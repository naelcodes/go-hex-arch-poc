package base

import "github.com/naelcodes/ab-backend/internal/common/types"

type IBaseEntity any

type BaseEntity struct {
	ID types.Id `json:"id"`
}

type GenericDTO[T any] any

type IApplicationService[T any] interface {
	Init() *IApplicationService[T]
	Execute(T) (T, error)
}

type IRepository[T any] interface {
	FindById(id types.Id) (*T, error)
	Find(query IQueryBuilder) ([]*T, error)
	FindAll() ([]*T, error)
	Count() int
	Save(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
}

type IQueryBuilder interface {
	build()
}
