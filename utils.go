package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func ParseQuery(queryVal interface{}) string {
	querystring := ""

	if queryVal == nil {
		return querystring
	}

	kind := reflect.TypeOf(queryVal).Kind()
	if kind == reflect.Struct {
		log.Printf("%v should be a value not struct", queryVal)
		return querystring
	} else if kind == reflect.Slice || kind == reflect.Array {
		arrVal := reflect.ValueOf(queryVal)
		for i := 0; i < arrVal.Len(); i++ {
			value := arrVal.Index(i).Interface()
			if querystring != "" {
				querystring += "&"
			}
			querystring += ParseQuery(value)
		}
	} else if kind == reflect.Map {
		mapVal := reflect.ValueOf(queryVal)
		for _, key := range mapVal.MapKeys() {
			value := mapVal.MapIndex(key).Interface()
			if querystring != "" {
				querystring += "&"
			}
			querystring += fmt.Sprintf("%s=%s", key, ParseQuery(value))
		}
	} else if kind == reflect.Ptr {
		ct := reflect.ValueOf(queryVal).Elem()
		for i := 0; i < ct.NumField(); i++ {
			field := ct.Type().Field(i)
			tagVal := field.Tag.Get("query")
			if tagVal == "" {
				continue
			}

			value := ParseQuery(ct.Field(i).Interface())
			if value != "" {
				if querystring != "" {
					querystring += "&"
				}
				querystring += fmt.Sprintf("%s=%s", tagVal, value)
			}
		}
	} else {
		switch v := queryVal.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			trans := strconv.Itoa(v.(int))
			querystring = trans
		case bool:
			trans := strconv.FormatBool(v)
			querystring = trans
		case string:
			querystring = v
		}
	}

	return querystring
}

func ParseData(data interface{}) io.Reader {
	if data == nil {
		return nil
	}
	buffer, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return strings.NewReader(string(buffer))
}

func IfSlashPrefixString(s string) string {
	if s == "" {
		return s
	}
	s = strings.TrimSuffix(s, "/")
	if strings.HasPrefix(s, "/") {
		return ToFormat(s)
	}
	return "/" + ToFormat(s)
}

// ToFormat takes a string and returns a formatted string. The string is
// converted to lowercase and spaces are removed.
func ToFormat(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
