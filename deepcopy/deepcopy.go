package deepcopy

import (
	"errors"
	"reflect"
)

func Copy(src, dst interface{}) error {
	return process(src, dst)
}

func process(src, dst interface{}) error {
	var (
		srcReflectValue = reflect.ValueOf(src)
		srcReflectType  = reflect.TypeOf(src)
		dstReflectValue = reflect.ValueOf(dst)
		dstReflectType  = reflect.TypeOf(dst)
	)

	for srcReflectType.Kind() == reflect.Ptr {
		srcReflectValue = srcReflectValue.Elem()
		srcReflectType = srcReflectValue.Type()
	}

	for dstReflectType.Kind() == reflect.Ptr {
		dstReflectValue = dstReflectValue.Elem()
		dstReflectType = dstReflectValue.Type()
	}

	if !dstReflectValue.CanSet() {
		return errors.New("type of dst must be pointer")
	}

	switch srcReflectType.Kind() {
	case reflect.Struct:
		for i := 0; i < srcReflectType.NumField(); i++ {
			srcFieldType := srcReflectType.Field(i)
			srcFieldValue := srcReflectValue.Field(i)

			dstField := dstReflectValue.FieldByName(srcFieldType.Name)
			if !dstField.IsValid() {
				continue
			}

			if srcFieldValue.Type() == dstField.Type() ||
				(dstField.Type().Kind() == reflect.Interface && dstField.NumMethod() == 0) {
				dstField.Set(reflect.ValueOf(srcFieldValue.Interface()))
			}
		}
	case reflect.Map:
	default:
		if srcReflectType.Kind() == dstReflectType.Kind() {
			dstReflectValue.Set(reflect.ValueOf(srcReflectValue.Interface()))
		}
	}

	return nil
}
