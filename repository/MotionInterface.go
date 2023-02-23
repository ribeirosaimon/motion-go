package repository

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type motionRepository[T any] interface {
	FindById(string) T
	FindAll() []T
	DeleteById(string) bool
	Save(T) T
	UpdateById(T) T
}

type motionStructRepository[T any] struct {
	myStruct T
}

func newMotionRepository[T any]() motionRepository[T] {
	var myStruct T
	return motionStructRepository[T]{
		myStruct: myStruct,
	}
}

func (m motionStructRepository[T]) FindById(s string) T {

	return m.myStruct
}

func createSqlField(t reflect.Type) string {
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
	elem := reflect.TypeOf(m.myStruct)

	var insertFieldNames string
	var insertFieldValues string

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if field.IsExported() {
			field.Name = strings.ToLower(field.Name[:1]) + field.Name[1:]
		}
		var fieldName string
		if value, ok := field.Tag.Lookup("json"); ok {
			fieldName = tagTreatment(&value)
		}
		fieldByName := reflect.ValueOf(structValue).Elem().FieldByName(field.Name)
		if fieldByName. {

		}
		insertFieldNames += fieldName
		insertFieldValues += fieldValue

	}
	insertSqlQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ",
		strings.ToLower(elem.Name()), insertFieldNames, insertFieldValues)

	fmt.Println(insertSqlQuery)
	return m.myStruct
}

func (m motionStructRepository[T]) UpdateById(t T) T {
	reflect.ValueOf(&t).Elem()
	panic("implement me")
}

func tagTreatment(json *string) string {
	stringNoOmitempty := strings.ReplaceAll(*json, "omitempty", "")
	re := regexp.MustCompile("[A-Z]")
	stringNoOmitempty = re.ReplaceAllStringFunc(stringNoOmitempty, func(match string) string {
		return "_" + strings.ToLower(match)
	})
	return stringNoOmitempty
}
