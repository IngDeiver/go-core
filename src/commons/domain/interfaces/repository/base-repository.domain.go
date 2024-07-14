package baseRepositoryDomain

import (
	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
	"go.mongodb.org/mongo-driver/bson"
)
 


type BaseRepositoryDomain[T any] interface {
	FindAll(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto, customPipeline bson.A ) (*dtos.PagedResponse[T], error)
	FindAllWithoutPagination(filter any, customPipeline bson.A ) ([]T, error)
	Create(document any) (T, error)
	UpdateById(ID string, document any) (*T, error)
	FindById(ID string) (T, error)
	RemoveById(ID string) (T, error)
	RemoveOne(filter interface{}) (*T, error)
	FindOne(filter interface{}) (*T, error)
	UpdateOne(filter interface{}, document any) (*T, error)
}
