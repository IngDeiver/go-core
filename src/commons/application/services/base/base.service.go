package baseMongoService

import (
	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
	baseRepoDomain "github.com/ingdeiver/go-core/src/commons/domain/interfaces/repository"
)

// implements BaseServiceDomain
type BaseService[T any] struct {
	Repository baseRepoDomain.BaseRepositoryDomain[T]
}

func  New[T any](repository baseRepoDomain.BaseRepositoryDomain[T])  BaseService[T] {
	return  BaseService[T]{repository}
}

func (s  *BaseService[T]) FindAll(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto ) (*dtos.PagedResponse[T], error){
	return s.Repository.FindAll(filter, pagination, sort)
}

func (s  *BaseService[T]) Create(data T) (T, error) {
	return s.Repository.Create(data)
}

func (s  *BaseService[T]) FindById(ID string) (T, error) {
	return s.Repository.FindById(ID)
}

func (s  *BaseService[T]) RemoveById(ID string) (T, error) {
	return s.Repository.RemoveById(ID)
}
