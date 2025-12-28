package fetch

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func BuildQueryParams(queryParam interface{}) url.Values {
	values := url.Values{}
	if queryParam == nil {
		return values
	}

	v := reflect.ValueOf(queryParam)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return values
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return values
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if fieldType.PkgPath != "" {
			continue
		}

		fieldName := fieldType.Name
		tag := fieldType.Tag.Get("query")
		if tag == "" {
			continue
		}
		if tag != "" {
			fieldName = strings.Split(tag, ",")[0]
		}

		val := reflect.Indirect(field)
		if !val.IsValid() {
			continue
		}

		switch val.Kind() {
		case reflect.Slice, reflect.Array:
			for j := 0; j < val.Len(); j++ {
				values.Add(fieldName, fmt.Sprintf("%v", val.Index(j).Interface()))
			}
		case reflect.Struct:
			if t, ok := val.Interface().(time.Time); ok {
				values.Add(fieldName, t.Format(time.RFC3339))
				continue
			}
			subValues := BuildQueryParams(val.Interface())
			for k, vls := range subValues {
				for _, v := range vls {
					values.Add(k, v)
				}
			}
		default:
			values.Add(fieldName, fmt.Sprintf("%v", val.Interface()))
		}
	}
	return values
}

func ParseQuery(params ...any) string {
	values := strings.Builder{}
	for _, param := range params {
		if param == nil {
			continue
		}
		v := BuildQueryParams(param)
		if values.Len() > 0 {
			values.WriteString("&")
		}
		values.WriteString(v.Encode())
	}

	return values.String()
}
