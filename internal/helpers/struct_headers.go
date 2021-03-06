package helpers

import (
	"reflect"
	"strings"
)

// Parse struct and return all the field names as headers
func GetStructHeaders(str interface{}) []string {
	v := reflect.ValueOf(str)
	t := v.Type()

	var out []string
	for i := 0; i < t.NumField(); i++ {
		out = append(out, strings.ToUpper(t.Field(i).Name))
	}

	return out
}
