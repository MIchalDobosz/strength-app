package utils

import (
	"errors"
	"reflect"
)

func StructTags(str any, tag string) ([]string, error) {
	typ := reflect.TypeOf(str)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("provided str is not a struct or pointer to a struct")
	}

	tags := []string{}
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get(tag)
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	return tags, nil
}

func SetStructField(str any, name string, value any) error {
	typ := reflect.TypeOf(str)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return errors.New("provided str is not a struct or pointer to a struct")
	}

	valueType := reflect.TypeOf(value)
	val := reflect.ValueOf(str)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Type().Name() == name {
			if field.Type() != valueType {
				return errors.New("provided value is not the same type as struct field")
			}
			if !field.CanSet() {
				return errors.New("struct field cannot be set")
			}
			field.Set(reflect.ValueOf(value))
			return nil
		}
	}
	return errors.New("struct does not contain field with provided name")
}
