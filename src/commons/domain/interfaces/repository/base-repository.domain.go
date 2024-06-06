package baseRepositoryDomain

import "go.mongodb.org/mongo-driver/bson"

type BaseRepositoryDomain[T any] interface {
	List(filter bson.D) ([]T, error)
	Add(document T) (T, error)
	Get(ID string) (T, error)
	Remove(ID string) (T, error)
}
