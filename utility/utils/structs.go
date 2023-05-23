package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// StructToStruct 使用反射，转换结构体
func StructToStruct(sourceStruct interface{}, targetStruct interface{}) {
	source := structToMap(sourceStruct)
	targetV := reflect.ValueOf(targetStruct)
	targetT := reflect.TypeOf(targetStruct)
	if targetV.Kind() == reflect.Ptr {
		targetV = targetV.Elem()
		targetT = targetT.Elem()
	}
	for i := 0; i < targetV.NumField(); i++ {
		fieldName := targetT.Field(i).Name
		sourceVal := source[fieldName]
		if !sourceVal.IsValid() {
			continue
		}
		targetVal := targetV.Field(i)
		fmt.Printf("targetVal.Type(): %+v; sourceVal.Type(): %+v\n", targetVal.Type(), sourceVal.Type())
		if targetVal.Type() != sourceVal.Type() {
			continue
		}
		targetVal.Set(sourceVal)
	}
}

func structToMap(structName interface{}) map[string]reflect.Value {
	t := reflect.TypeOf(structName)
	v := reflect.ValueOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	fieldNum := t.NumField()
	resMap := make(map[string]reflect.Value, fieldNum)
	for i := 0; i < fieldNum; i++ {
		resMap[t.Field(i).Name] = v.Field(i)
	}
	return resMap
}

func ArrayToString(array interface{}, sep string) (res string) {
	n := strings.Trim(fmt.Sprintf(gconv.String(array)), "[]")
	res = strings.Replace(n, " ", sep, -1)
	return
}

func StructToMap(source interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	t := reflect.TypeOf(source)
	v := reflect.ValueOf(source)
	fields := t.NumField()
	var matchAll = regexp.MustCompile("([a-z0-9])([A-Z])")
	for i := 0; i < fields; i++ {
		name := t.Field(i).Name
		newName := matchAll.ReplaceAllString(name, "${1}_${2}") //拆分单词
		newMap[gstr.ToLower(newName)] = v.Field(i).Interface()
	}
	return newMap
}
