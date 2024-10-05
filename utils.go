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

func ParseQuery(arrVal []interface{}) string {
	querystring := ""
	for _, val := range arrVal {
		if reflect.TypeOf(val).Kind() == reflect.Struct {
			log.Fatalf("%v should be a value not struct", val)
			continue
		}

		ct := reflect.ValueOf(val).Elem()
		for i := 0; i < ct.NumField(); i++ {
			field := ct.Type().Field(i)
			tagVal := field.Tag.Get("query")
			if tagVal == "" {
				continue
			}

			value := ct.Field(i).Interface()

			switch v := value.(type) {
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				trans := strconv.Itoa(v.(int))
				value = trans
			case bool:
				trans := strconv.FormatBool(v)
				value = trans
			}

			if value != "" {
				if querystring != "" {
					querystring += "&"
				}
				querystring += fmt.Sprintf("%s=%s", tagVal, value)
			}
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
