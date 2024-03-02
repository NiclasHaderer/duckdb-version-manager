package utils

import (
	"reflect"
)

func Keys(m any) []string {
	val := reflect.ValueOf(m)
	keys := make([]string, 0, val.Len())
	for _, key := range val.MapKeys() {
		keys = append(keys, key.String())
	}

	return keys
}
