package common

type IRepository[T any] interface {
	FindById(id uint) (*T, error)
	Find(query GetQueryParams, args ...any) ([]*T, error)
	FindAll() ([]*T, error)
	Count() int
	Save(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
}
