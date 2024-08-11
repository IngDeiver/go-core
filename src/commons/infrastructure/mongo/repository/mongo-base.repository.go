package commonMongoRepository

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/ingdeiver/go-core/src/commons/application/helpers"
	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
	errorsDomain "github.com/ingdeiver/go-core/src/commons/domain/errors"
	baseSchema "github.com/ingdeiver/go-core/src/commons/domain/interfaces"
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	"github.com/ingdeiver/go-core/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var l = logger.Get()

// implements BaseRepositoryDomain
type MongoBaseRepository[T any] struct {
	regexFields []string
}

func New[T any]() MongoBaseRepository[T] {
	return MongoBaseRepository[T]{regexFields: []string{}}
}

func (s *MongoBaseRepository[T]) GetRegexFields() []string {
	return s.regexFields
}

func (s *MongoBaseRepository[T]) SetRegexFields(regexFields []string){
	 s.regexFields = regexFields
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

func (s MongoBaseRepository[T]) buildFilter(filter any, regexFields []string) bson.D {
	v := reflect.ValueOf(filter)
	filterBson := bson.D{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i)

		// Check if the field has a "form" tag
		tag := field.Tag.Get("form")

		// Only include non-zero values in the filter
		if tag != "" && tag != "-" && value.IsValid() && !value.IsZero() {
			// Handle pointers
			if value.Kind() == reflect.Ptr && !value.IsNil() {
				value = value.Elem()
			}

			filterValue :=  value.Interface()

			// multiple values to filter
			if value.Kind() == reflect.String {
				strValue := value.Interface().(string)
				if strings.Contains(strValue, ",") {
					query := bson.M{
						"$in": strings.Split(strValue, ","),
					}
					filterValue = query
				}
			}

			// regex filter  
			if value.Kind() == reflect.String {
				fieldName := tag
				if helpers.ContainsStr(regexFields, fieldName){
					strValue := value.Interface().(string)
					pattern :=  fmt.Sprintf(".*%s.*", strValue)
					filterValue = bson.M{"$regex": pattern}
					
				}
			}

			filterBson = append(filterBson, bson.E{Key: tag, Value: filterValue})
		}
	}

	return filterBson
}

func (s MongoBaseRepository[T]) FindAll(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto, customPipeline bson.A) (*dtos.PagedResponse[T], error) {
    collectionName := s.getCollectionName()
    collection := config.GetCollection(collectionName)

    bsonFilter := s.buildFilter(filter, s.regexFields)
    pipeline := bson.A{}

    // set sort
    if sort != nil && sort.SortField != "" {
        direction := -1
        if sort.SortDirection == "asc" {
            direction = 1
        }
        pipeline = append(pipeline, bson.D{{Key: "$sort", Value: bson.D{{Key: sort.SortField, Value: direction}}}})
    }

    // set filters
    pipeline = append(pipeline, bson.D{{Key: "$match", Value: bsonFilter}})

    // include custom pipeline
    if len(customPipeline) > 0 {
        pipeline = append(pipeline, customPipeline...)
    }

    // create a pipeline for counting the total number of documents
    countPipeline := append(bson.A{
        bson.D{{Key: "$match", Value: bsonFilter}},
    }, customPipeline...)

    // perform the count
    countPipeline = append(countPipeline, bson.D{{Key: "$count", Value: "total"}})
    countCursor, err := collection.Aggregate(context.Background(), countPipeline)
    if err != nil {
        l.Error().Err(err).Msg("Error while counting documents")
        return nil, err
    }
    var countResult []struct {
        Total int `bson:"total"`
    }
    if err = countCursor.All(context.Background(), &countResult); err != nil {
        l.Error().Err(err).Msg("Error while fetching count result")
        return nil, err
    }
    totalCount := int64(0)
    if len(countResult) > 0 {
        totalCount = int64(countResult[0].Total)
    }
    
    // paginate data
    limit := int64(pagination.GetLimit())
    skip := int64((pagination.GetPage() - 1) * pagination.GetLimit())
    pipeline = append(pipeline, bson.D{{Key: "$skip", Value: skip}})
    pipeline = append(pipeline, bson.D{{Key: "$limit", Value: limit}})

    cur, err := collection.Aggregate(context.Background(), pipeline)
    if err != nil {
        l.Error().Err(err).Msg("Error while fetching documents")
        return nil, err
    }
    defer cur.Close(context.Background())

    var results []T
    if err = cur.All(context.Background(), &results); err != nil {
        l.Error().Err(err).Msg("Error while decoding documents")
        return nil, err
    }

    if results == nil {
        results = []T{}
    }

    totalPages := int((totalCount + limit - 1) / limit)

    pagedResponse := dtos.PagedResponse[T]{
        Data: results,
        PaginationMetadata: dtos.PaginationMetadata{
            Page:       pagination.GetPage(),
            Limit:      pagination.GetLimit(),
            TotalPages: totalPages,
            TotalCount: int(totalCount),
        },
    }

    return &pagedResponse, nil
}

func (s MongoBaseRepository[T]) FindAllWithoutPagination(filter any, customPipeline bson.A) ([]T, error) {
    collectionName := s.getCollectionName()
    collection := config.GetCollection(collectionName)

    bsonFilter := s.buildFilter(filter, s.regexFields)
    pipeline := bson.A{bson.D{{Key: "$match", Value: bsonFilter}}}

    // include custom pipeline
    if len(customPipeline) > 0 {
        pipeline = append(pipeline, customPipeline...)
    }

    
 
    cur, err := collection.Aggregate(context.Background(), pipeline)
    if err != nil {
        l.Error().Err(err).Msg("Error while fetching documents")
        return nil, err
    }
    defer cur.Close(context.Background())

    var results []T
    if err = cur.All(context.Background(), &results); err != nil {
        l.Error().Err(err).Msg("Error while decoding documents")
        return nil, err
    }

    if results == nil {
        results = []T{}
    }
    return results, nil
}

func (s MongoBaseRepository[T]) Create(user any) (T, error) {
	var result T
	collectionName := s.getCollectionName()
	collection := config.GetCollection(collectionName)

	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return result, err
	}

	err = helpers.SetFieldByReflection(&result, user, insertResult.InsertedID.(primitive.ObjectID), "ID")
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s MongoBaseRepository[T]) UpdateOne(filter interface{}, document any) (*T, error) {
	var result T
	collectionName := s.getCollectionName()
	collection := config.GetCollection(collectionName)

	body := bson.M{"$set": document}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := collection.FindOneAndUpdate(context.Background(), filter, body, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errorsDomain.ErrNotFoundError
		}
		return nil, err
	}

	return &result, nil
}

func (s MongoBaseRepository[T]) FindById(ID string) (T, error) {
	var result T
	collectionName := s.getCollectionName()
	collection := config.GetCollection(collectionName)
	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": objID}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, errorsDomain.ErrNotFoundError
		}
		return result, err
	}
	return result, nil
}

func (s MongoBaseRepository[T]) RemoveById(ID string) (T, error) {
	var result T
	collectionName := s.getCollectionName()
	collection := config.GetCollection(collectionName)
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return result, err
	}

	filter := bson.M{"_id": objID}
	err = collection.FindOneAndDelete(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, errorsDomain.ErrNotFoundError
		}
		return result, err
	}
	return result, nil
}

func (s *MongoBaseRepository[T]) UpdateById(ID string, document any) (*T, error) {
	var result T
	collectionName := s.getCollectionName()
	collection := config.GetCollection(collectionName)
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$set": document}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errorsDomain.ErrNotFoundError
		}
		return nil, err
	}
	return &result, nil
}

func (s MongoBaseRepository[T]) FindOne(filter interface{}) (*T, error) {
	var result T
	collectionName := s.getCollectionName()
	collection := config.GetCollection(collectionName)
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errorsDomain.ErrNotFoundError
		}
		return &result, err
	}
	return &result, nil
}

func (s MongoBaseRepository[T]) RemoveOne(filter interface{}) (*T, error) {
	var result T
	collectionName := s.getCollectionName()
	collection := config.GetCollection(collectionName)
	err := collection.FindOneAndDelete(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errorsDomain.ErrNotFoundError
		}
		return &result, err
	}
	return &result, nil
}
