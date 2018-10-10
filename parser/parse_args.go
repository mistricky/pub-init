package parser

import (
	"../error_checker"
	"errors"
	"reflect"
)

func Slice(args interface{}) []interface{} {
	slice, err := parseArgs(args, reflect.Array)
	checker.CheckErr(err)
	length := slice.Len()
	val := make([]interface{}, length)

	for i := 0;i < length;i++ {
		val[i] = slice.Index(i).Interface()
	}

	return val
}

func parseArgs(args interface{}, kind reflect.Kind) (reflect.Value, error){
	val := reflect.ValueOf(args)

	if val.Kind() == kind {
		return val, errors.New("[function]parseArgs:type error")
	} else {
		return val, nil
	}
}