package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/ribeirosaimon/motion-go/pkg/config/database"
)

type motionRepository[T any] interface {
	FindById(string) T
	FindAll() []T
	DeleteById(string) bool
	Save(T) T
	UpdateById(T) T
}

type motionStructRepository[T any] struct {
	myStruct   T
	connection *sql.DB
}

func newMotionRepository[T any]() motionRepository[T] {
	var myStruct T
	connect, err := database.Connect()
	if err != nil {
		panic(err)
	}
	return motionStructRepository[T]{
		myStruct:   myStruct,
		connection: connect,
	}
}

func (m motionStructRepository[T]) FindById(s string) T {

	return m.myStruct
}

func convertFieldToValue(t reflect.Type) string {
	kind := t.Kind()
	switch kind {
	case reflect.Uint, reflect.Int, reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16, reflect.Uint32, reflect.Int32:
		return "int"
	case reflect.Uint64, reflect.Int64:
		return "bigint"
	case reflect.Bool:
		return "boolean"
	case reflect.String:
		return "string"
	default:
		return "string"
	}
}

func (m motionStructRepository[T]) FindAll() []T {
	panic("implement me")
}

func (m motionStructRepository[T]) DeleteById(s string) bool {
	return false
}

func (m motionStructRepository[T]) Save(structValue T) T {
	defer m.connection.Close()
	reflectValueOf := reflect.ValueOf(m.myStruct)
	reflectTypeOf := reflectValueOf.Type()

	var queryStringName, queryStringValue string

	for i := 0; i < reflectTypeOf.NumField(); i++ {
		var fieldName string
		field := reflectTypeOf.Field(i)

		if field.IsExported() {
			fieldName = strings.ToLower(field.Name[:1]) + field.Name[1:]
		}

		if value, ok := field.Tag.Lookup("json"); ok {
			fieldName = tagTreatment(&value)
		}

		structToSaveReflections := reflect.ValueOf(structValue)
		fieldValue := structToSaveReflections.FieldByName(reflectTypeOf.Field(i).Name)

		defaultValue := reflect.Zero(fieldValue.Type())

		if !reflect.DeepEqual(fieldValue.Interface(), defaultValue.Interface()) && !fieldValue.IsZero() {
			queryStringName += fmt.Sprintf("%s,", fieldName)
			queryStringValue += fmt.Sprint(fieldValue, ",")
		}

	}
	queryStringName = strings.TrimSuffix(queryStringName, ",")
	queryStringValue = strings.TrimSuffix(queryStringValue, ",")
	insertSqlQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		strings.ToLower(reflectTypeOf.Name()), queryStringName, queryStringValue)

	_, err := m.connection.Exec(insertSqlQuery)
	if err != nil {
		fmt.Errorf("error in execute query")
	}

	return m.myStruct
}

func (m motionStructRepository[T]) UpdateById(t T) T {
	reflect.ValueOf(&t).Elem()
	panic("implement me")
}

func tagTreatment(json *string) string {
	stringNoOmitempty := strings.ReplaceAll(*json, "omitempty", "")
	stringNoOmitempty = strings.ReplaceAll(stringNoOmitempty, ",", "")
	re := regexp.MustCompile("[A-Z]")
	stringNoOmitempty = re.ReplaceAllStringFunc(stringNoOmitempty, func(match string) string {
		return "_" + strings.ToLower(match)
	})
	return stringNoOmitempty
}
