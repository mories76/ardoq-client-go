package ardoq

import (
	"fmt"
	"reflect"
)

// this function uses reflection and type assertion
// to check and remove nil values
func removeNulls(m map[string]interface{}) {
	val := reflect.ValueOf(m)
	for _, e := range val.MapKeys() {
		v := val.MapIndex(e)
		if v.IsNil() {
			delete(m, e.String())
			continue
		}

		switch t := v.Interface().(type) {
		// If key is a JSON object (Go Map), use recursion to go deeper
		case map[string]interface{}:
			removeNulls(t)
		}
	}
}

// typeConversion deals with the fact that ardoq know many different field types,
// the Terraform provider SDK v2(.7) does not have dynamic schema types
// so here conversion to string is done.
func typeConversion(m map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		fmt.Printf("conversion for %s from '%T'\n", k, v)
		// option 1
		switch v := v.(type) {
		// case string:
		// 	result[k] = v
		// case []string:
		// 	result[k] = strings.Join(v, ",")
		// case int:
		// 	result[k] = "type was int"
		default:
			fmt.Printf("lets try default conversion for %s from '%T'\n", k, v)
			result[k] = fmt.Sprint(v)
		}

		// // option 2
		// if v, ok := v.([]string); ok {
		// 	m[k] = v
		// }
	}

	return result
}
