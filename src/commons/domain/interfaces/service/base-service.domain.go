package baseServiceDomain

import (
	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
)

type BaseServiceDomain[T any] interface {
	List(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto) (*dtos.PagedResponse[T], error)
	Add(document T) (T, error)
	Get(ID string) (T, error)
	Remove(ID string) (T, error)
}
