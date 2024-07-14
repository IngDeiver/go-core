package baseServiceDomain

import (
	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
)

type BaseServiceDomain[T any] interface {
	FindAll(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto) (*dtos.PagedResponse[T], error)
	FindAllWithoutPagination(filter any) ([]T, error)
	Create(document any) (T, error)
	UpdateById(ID string, document any) (*T, error)
	FindById(ID string) (T, error)
	RemoveById(ID string) (T, error)
	RemoveOne(filter interface{}) (*T, error)
	FindOne(filter interface{}) (*T, error)
	UpdateOne(filter interface{}, document any) (*T, error)
}
