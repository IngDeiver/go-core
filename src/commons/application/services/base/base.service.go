package baseMongoService

import (
	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
	baseRepoDomain "github.com/ingdeiver/go-core/src/commons/domain/interfaces/repository"
	"go.mongodb.org/mongo-driver/bson"
)

// implements BaseServiceDomain
type BaseService[T any] struct {
	repository baseRepoDomain.BaseRepositoryDomain[T]
}

func  New[T any](repository baseRepoDomain.BaseRepositoryDomain[T])  BaseService[T] {
	return  BaseService[T]{repository}
}

func (s  *BaseService[T]) FindAll(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto) (*dtos.PagedResponse[T], error){
	return s.repository.FindAll(filter, pagination, sort, nil)
}

func (s  *BaseService[T]) FindAllWithoutPagination(filter any) ([]T, error){
	return s.repository.FindAllWithoutPagination(filter,bson.A{})
}

func (s  *BaseService[T]) Create(data any) (T, error) {
	return s.repository.Create(data)
}

func (s  *BaseService[T]) UpdateOne(filter interface{}, document any) (*T, error) {
	return s.repository.UpdateOne(filter, document)
}

func (s  *BaseService[T]) FindById(ID string) (T, error) {
	return s.repository.FindById(ID)
}

func (s  *BaseService[T]) RemoveById(ID string) (T, error) {
	return s.repository.RemoveById(ID)
}

func (s  *BaseService[T]) UpdateById(ID string, document any) (*T, error) {
	return s.repository.UpdateById(ID, document)
}

func (s  *BaseService[T]) FindOne(filter interface{}) (*T, error) {
	return s.repository.FindOne(filter)
}

func (s  *BaseService[T]) RemoveOne(filter interface{}) (*T, error) {
	return s.repository.RemoveOne(filter)
}

