package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MotionSQLRepository[T Entity] struct {
	myStruct Entity
	database *gorm.DB
	timeDone uint8
}

func newMotionSQLRepository[T Entity](gormConnection *gorm.DB) *MotionSQLRepository[T] {
	var myStruct T
	return &MotionSQLRepository[T]{
		myStruct: myStruct,
		database: gormConnection,
	}
}

func (m MotionSQLRepository[T]) ExistByField(field string, fieldvalue interface{}) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.timeDone))
	defer cancel()
	var value T
	tx := m.database.Where(fmt.Sprintf("%s = ?", field), fieldvalue).Find(&value)
	if tx.RowsAffected > 0 {
		ctx.Done()
		return true
	}
	ctx.Done()
	return false
}

func (m MotionSQLRepository[T]) FindWithPreloads(preloads string, s interface{}) (T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.timeDone))
	defer cancel()
	var value T
	tx := m.database.Preload(preloads).Find(&value, s)
	if tx.RowsAffected == 0 || tx.Error != nil {
		ctx.Done()
		return value, errors.New("values not found")
	}

	return value, nil
}

func (m MotionSQLRepository[T]) FindByField(field string, fieldvalue interface{}) (T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.timeDone))
	defer cancel()
	var value T
	tx := m.database.Where(fmt.Sprintf("%s = ?", field), fieldvalue).Find(&value)
	if tx.RowsAffected == 0 || tx.Error != nil {
		ctx.Done()
		return value, errors.New("values not found")
	}

	return value, nil
}

func (m MotionSQLRepository[T]) FindById(s interface{}) (T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.timeDone))
	defer cancel()
	var value T
	tx := m.database.Find(&value, s)
	if tx.RowsAffected == 0 || tx.Error != nil {
		ctx.Done()
		return value, errors.New("values not found")
	}

	return value, nil
}

func (m MotionSQLRepository[T]) FindAll(limit, page int) ([]T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.timeDone))
	defer cancel()
	var values []T
	tx := m.database.Limit(limit).Offset(page).Find(&values)
	if err := tx.Error; err != nil {
		ctx.Done()
		return nil, errors.New("values not found")
	}
	return values, nil
}

func (m MotionSQLRepository[T]) DeleteById(s interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.timeDone))
	defer cancel()
	value, err := m.FindById(s)
	if err != nil {
		ctx.Done()
		return err
	}
	tx := m.database.Delete(&value, s)
	if tx.Error == nil && tx.RowsAffected > 0 {
		return nil
	}
	return errors.New("values not found")
}

func (m MotionSQLRepository[T]) Save(structValue T) (T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.timeDone))
	defer cancel()
	var value T
	if err := m.database.AutoMigrate(&value); err != nil {
		ctx.Done()
		return value, fmt.Errorf(err.Error())
	}
	if err := m.database.Save(&structValue).Error; err != nil {
		ctx.Done()
		return value, fmt.Errorf(err.Error())
	}
	return m.FindById(structValue.GetId())

}
