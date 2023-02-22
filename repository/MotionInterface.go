package repository

import (
	"fmt"
	"reflect"
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
	elem := reflect.TypeOf(m.myStruct)
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if field.IsExported() {
			field.Name = strings.ToLower(field.Name[:1]) + field.Name[1:]
		}
		fmt.Println(field.Name, field.Type)
	}
	return m.myStruct
}

func (m motionStructRepository[T]) FindAll() []T {
	// TODO implement me
	panic("implement me")
}

func (m motionStructRepository[T]) DeleteById(s string) bool {
	// TODO implement me
	panic("implement me")
}

func (m motionStructRepository[T]) Save(t T) T {
	// TODO implement me
	panic("implement me")
}

func (m motionStructRepository[T]) UpdateById(t T) T {
	reflect.ValueOf(&t).Elem()
	panic("implement me")
}
