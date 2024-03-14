package commonMongoService

import (
	commonMongoRepoDomain "github.com/ingdeiver/go-core/src/commons/domain/repository"
	commonMongoRepository "github.com/ingdeiver/go-core/src/commons/infrastructure/mongo/repository"
)

// implements ServiceDomain
type BaseMongoService[T any] struct {
	Repository commonMongoRepoDomain.BaseRepositoryDomain[T]
}

func  New[T any]()  BaseMongoService[T] {
	Repository := commonMongoRepository.New[T]()
	return  BaseMongoService[T]{Repository}
}

func (s  *BaseMongoService[T]) List() ([]T, error) {
	return s.Repository.List()
}

func (s  *BaseMongoService[T]) Add(data T) (T, error) {
	return s.Repository.Add(data)
}

func (s  *BaseMongoService[T]) Get(ID string) (T, error) {
	return s.Repository.Get(ID)
}

func (s  *BaseMongoService[T]) Remove(ID string) (T, error) {
	return s.Repository.Remove(ID)
}
