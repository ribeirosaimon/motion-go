package repository

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type motionNoSQLRepository[T Entity] struct {
	collection *mongo.Collection
	context    context.Context
}

func newMotionNoSQLRepository[T Entity](mongoConnection *mongo.Client) *motionNoSQLRepository[T] {
	var myStruct T
	collection := mongoConnection.Database("motion-go").Collection(reflect.TypeOf(myStruct).Name())

	return &motionNoSQLRepository[T]{
		collection: collection,
		context:    context.Background(),
	}
}

func (m motionNoSQLRepository[T]) FindById(i interface{}) (T, error) {
	return m.FindByField("_id", i)
}

func (m motionNoSQLRepository[T]) FindByField(s string, i interface{}) (T, error) {
	var myStruct T
	filter := bson.D{primitive.E{Key: s, Value: i}}
	err := m.collection.FindOne(m.context, filter).Decode(&myStruct)
	if err != nil {
		return myStruct, err
	}
	return myStruct, nil
}

func (m motionNoSQLRepository[T]) ExistByField(s string, i interface{}) bool {
	filter := bson.D{primitive.E{Key: s, Value: i}}
	count, err := m.collection.CountDocuments(m.context, filter)
	if err != nil || count == 0 {
		return false
	}
	return true
}

func (m motionNoSQLRepository[T]) FindAll(limit, page int) ([]T, error) {
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

func (m motionNoSQLRepository[T]) DeleteById(i interface{}) error {
	filter := bson.D{primitive.E{Key: "_id", Value: i}}
	_, err := m.collection.DeleteOne(m.context, filter)
	if err != nil {
		return err
	}
	return nil
}

func (m motionNoSQLRepository[T]) Save(t T) (T, error) {
	var myStruct T
	save, err := m.collection.InsertOne(m.context, t)
	if err != nil {
		return myStruct, err
	}
	return m.FindById(save.InsertedID)

}
