package repository

type Entity interface {
	GetId() interface{}
}

type MotionRepository[T Entity] interface {
	FindById(interface{}) (T, error)
	FindByField(string, interface{}) (T, error)
	ExistByField(string, interface{}) bool
	FindAll(int, int) ([]T, error)
	DeleteById(interface{}) error
	Save(T) (T, error)
}
