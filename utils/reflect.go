package utils

import "reflect"

func ReflectFields(dst, src interface{}) {
	dstVal := reflect.ValueOf(dst).Elem()
	srcVal := reflect.ValueOf(src).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		dstField := dstVal.Field(i)

		if !isEmptyValue(srcField) {
			dstField.Set(srcField)
		}
	}
}

func isEmptyValue(v reflect.Value) bool {
	return v.IsZero()
}
