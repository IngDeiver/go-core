package commonMongoRepository

import (
	"context"
	"reflect"
	"strings"

	baseSchema "github.com/ingdeiver/go-core/src/commons/domain/interfaces"
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	"github.com/ingdeiver/go-core/src/config"
	"go.mongodb.org/mongo-driver/bson"
)

var l = logger.Get()

// implements BaseRepositoryDomain
type MongoBaseRepository[T any] struct {
	
}

func  New[T any]()  MongoBaseRepository[T] {
	return  MongoBaseRepository[T]{}
}

func (s MongoBaseRepository[T]) getCollectionName() string {
	var t T
	if namer, ok := any(t).(baseSchema.CollectionNamer); ok {
		return namer.CollectionName()
	}
	// Si el tipo no implementa CollectionNamer, usamos reflection
	typ := reflect.TypeOf(t)
	name := typ.Name()
	return strings.ToLower(name) + "s" // Por ejemplo, "User" se convierte en "users"

}


func (s  MongoBaseRepository[T]) List(filter bson.D) ([]T, error) {
	collectionName := s.getCollectionName()
	collection := config.GetCollection(collectionName)
	cur, err := collection.Find(context.Background(), filter)
	if err != nil { 
		l.Error().Err(err)
		return []T{}, err
	 }
	defer cur.Close(context.Background())

	var results []T
	if err = cur.All(context.Background(), &results); err != nil {
		l.Error().Err(err)
		return  []T{}, err
	}

	return results, nil
}

func (s  MongoBaseRepository[T]) Add(user T) (T, error) {
	var result T
	return result, nil
}
func (s  MongoBaseRepository[T]) Get(ID string) (T, error) {
	var result T
	return result, nil
}
func (s  MongoBaseRepository[T]) Remove(ID string) (T, error) {
	var result T
	return result, nil
}
