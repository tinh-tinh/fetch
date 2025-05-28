package fetch

import (
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/tinh-tinh/tinhtinh/v2/core"
)

// ParseQuery takes an interface{} argument and returns a string representing
// the query string form of the argument. The argument can be a struct, a slice,
// an array, a map, a pointer, or a primitive type. The function will recursively
// traverse the argument and construct the query string accordingly.
//
// For structs, the function will extract the fields that have the "query" tag
// and use the tag value as the key in the query string. The value of the field
// will be converted to a string using the fmt package.
//
// For slices and arrays, the function will iterate over the elements and
// construct the query string by concatenating the string representation of each
// element with "&".
//
// For maps, the function will iterate over the key-value pairs and construct the
// query string by concatenating the key-value pairs with "&".
//
// For pointers, the function will recursively call itself on the value that the
// pointer points to.
//
// For primitive types, the function will use the fmt package to convert the
// value to a string.
//
// If the argument is nil, the function will return an empty string.
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

// ParseData serializes the given data into JSON format and returns an io.Reader
// containing the JSON data. If the input data is nil, the function returns nil.
// If the serialization fails, the function panics. The returned io.Reader can
// be used to read the JSON data as a stream.
func ParseData(data interface{}, encoder core.Encode) io.Reader {
	if data == nil {
		return nil
	}
	buffer, err := encoder(data)
	if err != nil {
		panic(err)
	}
	return strings.NewReader(string(buffer))
}
