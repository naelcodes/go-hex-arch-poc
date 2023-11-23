package common

type IRepository[T any] interface {
	GetById(id uint) (*T, error)
	Save(entity *T) error
}

type IApplication interface {
	init(GlobalContext)
}
