package order

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Order(data interface{}) string {
	text, _ := json.Marshal(data)
	temp := make(map[string]interface{})
	_ = json.Unmarshal(text, &temp)
	Filter(temp)
	result, _ := json.Marshal(temp)
	return string(result)
}

func Filter(data map[string]interface{}) {
	if _, ok := data["sign"]; ok {
		delete(data, "sign")
	}
	for _, value := range data {
		if elem, ok := value.(map[string]interface{}); ok {
			Filter(elem)
		}
	}
}
func serilizationSign(data interface{}) string {
	result := "{"
	typeOf := reflect.TypeOf(data)
	slice := make([]string, 0, typeOf.NumField())
	for i := 0; i < typeOf.NumField(); i++ {
		name := reflect.TypeOf(data).Field(i).Tag.Get("json")
		name = strings.Split(name, ",")[0]
		if name == "sign" {
			continue
		}
		t := reflect.TypeOf(data).Field(i).Type.Kind().String()
		value := fmt.Sprintf("%v", reflect.ValueOf(data).Field(i))
		if value == "" {
			continue
		}
		if reflect.ValueOf(data).Field(i).MethodByName("String").IsValid() {
			value = "\""
			value += fmt.Sprintf("%v", reflect.ValueOf(data).Field(i))
			value += "\""
		}
		if t == "string" {
			value = "\""
			value += fmt.Sprintf("%v", reflect.ValueOf(data).Field(i))
			value += "\""
		}
		if t == "struct" {
			value = serilizationSign(reflect.ValueOf(data).Field(i).Interface())
		}
		name = "\"" + name
		name += "\":" + value
		slice = append(slice, name)
	}
	sort.Strings(slice)
	result += strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ",", -1) + "}"
	return result
}

