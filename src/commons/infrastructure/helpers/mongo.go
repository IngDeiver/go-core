package helpers

import (
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CopyAndSetID copies the data from user to result and assigns the _id field
func CopyAndSetID[T any](result *T, domain any, id primitive.ObjectID) error {
    resultVal := reflect.ValueOf(result).Elem()
    userVal := reflect.ValueOf(domain)

    // Copy the data from domain to result
    if userVal.Kind() == reflect.Ptr {
        userVal = userVal.Elem()
    }

    for i := 0; i < userVal.NumField(); i++ {
        resultVal.Field(i).Set(userVal.Field(i))
    }

    // Assign the _id field to result
    idField := resultVal.FieldByName("ID")
    if !idField.IsValid() || !idField.CanSet() {
        return errors.New("ID field not found or cannot be set")
    }

    idField.Set(reflect.ValueOf(id))

    return nil
}