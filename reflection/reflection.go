package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfFields := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Struct:
		numberOfFields = val.NumField()
		getField = val.Field
	case reflect.Slice:
		numberOfFields = val.Len()
		getField = val.Index
	case reflect.String:
		fn(val.String())
	}

	for i := 0; i < numberOfFields; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}

	return val
}
