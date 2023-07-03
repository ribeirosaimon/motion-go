package repository

import (
	"context"
	"errors"
	"reflect"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MotionNoSQLRepository[T Entity] struct {
	collection *mongo.Collection
	context    context.Context
}

func newMotionNoSQLRepository[T Entity](ctx context.Context, mongoConnection *mongo.Client) *MotionNoSQLRepository[T] {
	var myStruct T
	collection := mongoConnection.Database(db.Conn.DatabaseName).Collection(reflect.TypeOf(myStruct).Name())
	return &MotionNoSQLRepository[T]{
		collection: collection,
		context:    ctx,
	}
}

func (m MotionNoSQLRepository[T]) FindById(i interface{}) (T, error) {
	var myStruct T

	id, ok := i.(string)
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return myStruct, err
	}
	if !ok {
		return myStruct, errors.New("id need to be a string")
	}

	return m.FindByField("_id", hex)
}

func (m MotionNoSQLRepository[T]) FindByField(s string, i interface{}) (T, error) {
	var myStruct T
	filter := bson.D{primitive.E{Key: s, Value: i}}
	err := m.collection.FindOne(m.context, filter).Decode(&myStruct)
	if err != nil {
		return myStruct, err
	}
	return myStruct, nil
}

func (m MotionNoSQLRepository[T]) ExistByField(s string, i interface{}) bool {
	filter := bson.D{primitive.E{Key: s, Value: i}}
	count, err := m.collection.CountDocuments(m.context, filter)
	if err != nil || count == 0 {
		return false
	}
	return true
}

func (m MotionNoSQLRepository[T]) FindAll(limit, page int) ([]T, error) {
	options := options.Find().SetSkip(int64(page)).SetLimit(int64(limit))

	cursor, err := m.collection.Find(m.context, bson.M{}, options)

	if err != nil {
		return nil, err
	}
	var results []T
	if err := cursor.All(m.context, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (m MotionNoSQLRepository[T]) DeleteById(i interface{}) error {
	filter := bson.D{primitive.E{Key: "_id", Value: i}}
	_, err := m.collection.DeleteOne(m.context, filter)
	if err != nil {
		return err
	}
	return nil
}

func (m MotionNoSQLRepository[T]) Save(t T) (T, error) {
	id, ok := t.GetId().(string)
	hex, err := primitive.ObjectIDFromHex(id)
	if !ok {
		return t, errors.New("id need to be a string")
	}

	filter := bson.D{primitive.E{Key: "_id", Value: hex}}
	countDocs, err := m.collection.CountDocuments(m.context, filter)
	if err != nil || countDocs == 0 {
		save, err := m.collection.InsertOne(m.context, t)
		insertedID, ok := save.InsertedID.(primitive.ObjectID)

		if !ok || err != nil {
			return t, errors.New("id need to be a string")
		}

		return m.FindById(insertedID.Hex())
	}
	_, err = m.collection.ReplaceOne(m.context, filter, t)
	if err != nil {
		return t, err
	}
	return m.FindById(hex.Hex())
}

func (m MotionNoSQLRepository[T]) GetCollection() *mongo.Collection {
	return m.collection
}
