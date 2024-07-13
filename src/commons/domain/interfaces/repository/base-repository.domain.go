package baseRepositoryDomain

import "github.com/ingdeiver/go-core/src/commons/domain/dtos"
 


type BaseRepositoryDomain[T any] interface {
	FindAll(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto ) (*dtos.PagedResponse[T], error)
	Create(document T) (T, error)
	FindById(ID string) (T, error)
	RemoveById(ID string) (T, error)
}
