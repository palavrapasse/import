package database

import (
	"reflect"
	"strings"

	"github.com/palavrapasse/import/internal/entity"
)

type Record interface{}
type Records []Record

type Field string
type Value any

func Fields(r Record) []Field {
	rf := reflectFields(r)
	fs := make([]Field, len(rf))

	for i, f := range rf {
		fs[i] = Field(strings.ToLower(f.Name))
	}

	return fs
}

func Values(r Record) []Value {
	rf := reflectFields(r)
	rv := reflect.ValueOf(r)
	vs := make([]Value, len(rf))

	for i, f := range rf {
		vs[i] = rv.FieldByName(f.Name).Interface()
	}

	return vs
}

func CopyWithNewKey(r Record, k entity.AutoGenKey) Record {
	copy := reflect.ValueOf(r).MethodByName("Copy")

	if !copy.IsZero() {
		return copy.Call([]reflect.Value{reflect.ValueOf(k)})[0].Interface()
	} else {
		return r
	}
}

func reflectFields(r Record) []reflect.StructField {
	return reflect.VisibleFields(reflect.TypeOf(r))
}
