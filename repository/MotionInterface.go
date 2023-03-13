package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type Entity interface {
	GetId() interface{}
}

type MotionRepository[T Entity] interface {
	FindById(interface{}) (T, error)
	FindByField(string, interface{}) (T, error)
	FindAll(int, int) ([]T, error)
	DeleteById(interface{}) error
	FindWithPreloads(string, interface{}) (T, error)
	Save(T) (T, error)
}

type motionStructRepository[T Entity] struct {
	myStruct Entity
	database *gorm.DB
}

func newMotionRepository[T Entity](gormConnection *gorm.DB) MotionRepository[T] {
	var myStruct T
	return motionStructRepository[T]{
		myStruct: myStruct,
		database: gormConnection,
	}
}

func (m motionStructRepository[T]) FindWithPreloads(preloads string, s interface{}) (T, error) {
	var value T
	tx := m.database.Preload(preloads).Find(&value, s)
	if tx.RowsAffected == 0 || tx.Error != nil {
		return value, fmt.Errorf(tx.Error.Error())
	}

	return value, nil
}

func (m motionStructRepository[T]) FindByField(field string, fieldvalue interface{}) (T, error) {
	var value T
	tx := m.database.Where(fmt.Sprintf("%s = ?", field), fieldvalue).Find(&value)
	if tx.RowsAffected == 0 || tx.Error != nil {
		return value, fmt.Errorf(tx.Error.Error())
	}

	return value, nil
}

func (m motionStructRepository[T]) FindById(s interface{}) (T, error) {
	var value T
	tx := m.database.Find(&value, s)
	if tx.RowsAffected == 0 || tx.Error != nil {
		return value, fmt.Errorf(tx.Error.Error())
	}

	return value, nil
}

func (m motionStructRepository[T]) FindAll(limit, page int) ([]T, error) {
	var values []T
	tx := m.database.Limit(limit).Offset(page).Find(&values)
	if err := tx.Error; err != nil {
		return nil, fmt.Errorf(tx.Error.Error())
	}
	return values, nil
}

func (m motionStructRepository[T]) DeleteById(s interface{}) error {
	value, err := m.FindById(s)
	if err != nil {
		return err
	}
	tx := m.database.Delete(&value, s)
	if tx.Error == nil && tx.RowsAffected > 0 {
		return nil
	}
	return fmt.Errorf(tx.Error.Error())
}

func (m motionStructRepository[T]) Save(structValue T) (T, error) {
	var value T
	if err := m.database.AutoMigrate(&value); err != nil {
		return value, fmt.Errorf(err.Error())
	}
	if err := m.database.Save(&structValue).Error; err != nil {
		return value, fmt.Errorf(err.Error())
	}
	return m.FindById(structValue.GetId())

}
