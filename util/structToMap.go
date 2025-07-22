package util

import "reflect"

func StructToMap(stc interface{}) map[string]interface{} {
	v := reflect.ValueOf(stc)
	// 判断反射对象是否为指针，指针需要取引用然后取值
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	m := make(map[string]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		// map的key值使用标签的值
		m[v.Type().Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	return m
}
