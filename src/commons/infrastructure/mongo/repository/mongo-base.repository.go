package commonMongoRepository

import (
	"context"
	"reflect"
	"strings"

	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
	baseSchema "github.com/ingdeiver/go-core/src/commons/domain/interfaces"
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	"github.com/ingdeiver/go-core/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s MongoBaseRepository[T]) buildFilter(filter any) bson.D {
	v := reflect.ValueOf(filter)
	filterBson := bson.D{}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i).Interface()
		tag := field.Tag.Get("json")

		if tag != "" && tag != "-" && !reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface()) {
			filterBson = append(filterBson, bson.E{Key: tag, Value: value})
		}
	}

	return filterBson
}

func (s  MongoBaseRepository[T]) List(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto ) (*dtos.PagedResponse[T], error) {
	collectionName := s.getCollectionName()
	collection := config.GetCollection(collectionName)

	// paginate data
	findOptions := options.Find()
	limit := int64(pagination.GetLimit())
	skip := int64((pagination.GetPage() - 1) * pagination.GetLimit())
	findOptions.SetLimit(limit)
	findOptions.SetSkip(skip)

	// Apply sort
	if sort != nil && sort.Field != "" {
		direction := 1
		if sort.Direction == "desc" {
			direction = -1
		}
		findOptions.SetSort(bson.D{{Key: sort.Field, Value: direction}})
	}

	bsonFilter := s.buildFilter(filter)
	cur, err := collection.Find(context.Background(), bsonFilter, findOptions)
	if err != nil { 
		l.Error().Err(err)
		return nil, err
	 }
	defer cur.Close(context.Background())

	var results []T
	if err = cur.All(context.Background(), &results); err != nil {
		l.Error().Err(err)
		return  nil, err
	}

	// if result is nil, so initialize an empty list
    if results == nil {
        results = []T{}
    }
	// Calculate total pages
	totalCount, err := collection.CountDocuments(context.Background(), bsonFilter)
	if err != nil {
		l.Error().Err(err)
		return nil, err
	}

	totalPages := int((totalCount + limit - 1) / limit)

	pagedResponse := dtos.PagedResponse[T]{
		Data: results,
		PaginationMetadata: dtos.PaginationMetadata{
			Page:       pagination.GetPage(),
			Limit:      pagination.GetLimit(),
			TotalPages: totalPages,
		},
	}

	return &pagedResponse, nil
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
