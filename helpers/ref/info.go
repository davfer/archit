package ref

import (
	"errors"
	"reflect"
)

var (
	ErrFieldNotFound  = errors.New("field not found")
	ErrCannotSetField = errors.New("cannot set field")
)

func GetGenericName[T any]() string {
	var z [0]T
	return reflect.TypeOf(z).Elem().Name()
}

func SetField[K any](v K, name string, value any) (K, error) {
	r := reflect.ValueOf(&v).Elem()
	f := r.FieldByName(name)
	if !f.IsValid() {
		return v, ErrFieldNotFound
	}
	if !f.CanSet() {
		return v, ErrCannotSetField
	}
	f.Set(reflect.ValueOf(value))
	return v, nil
}
