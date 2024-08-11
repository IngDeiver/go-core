package helpers

import (
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SetFieldByReflection copies the data from domain to result and assigns the field
func SetFieldByReflection[T any](result *T, domain any, id primitive.ObjectID, fieldName string) error {
    resultValue := reflect.ValueOf(result).Elem()
    reflectValue := reflect.ValueOf(domain)

    // Copy the data from domain to result
    if reflectValue.Kind() == reflect.Ptr {
        reflectValue = reflectValue.Elem()
    }

    for i := 0; i < reflectValue.NumField(); i++ {
        resultValue.Field(i).Set(reflectValue.Field(i))
    }

    // Assign the _id field to result
    idField := resultValue.FieldByName(fieldName)
    if !idField.IsValid() || !idField.CanSet() {
        return errors.New("ID field not found or cannot be set")
    }

    idField.Set(reflect.ValueOf(id))

    return nil
}