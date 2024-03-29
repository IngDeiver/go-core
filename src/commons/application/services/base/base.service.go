package baseMongoService

import (
	baseRepoDomain "github.com/ingdeiver/go-core/src/commons/domain/interfaces/repository"
)

// implements BaseServiceDomain
type BaseService[T any] struct {
	Repository baseRepoDomain.BaseRepositoryDomain[T]
}

func  New[T any](repository baseRepoDomain.BaseRepositoryDomain[T])  BaseService[T] {
	return  BaseService[T]{repository}
}

func (s  *BaseService[T]) List() ([]T, error) {
	return s.Repository.List()
}

func (s  *BaseService[T]) Add(data T) (T, error) {
	return s.Repository.Add(data)
}

func (s  *BaseService[T]) Get(ID string) (T, error) {
	return s.Repository.Get(ID)
}

func (s  *BaseService[T]) Remove(ID string) (T, error) {
	return s.Repository.Remove(ID)
}
