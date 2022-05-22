package fkmap

import "reflect"

// MapKeys is a slice of keys of a map.
func MapKeys(mapValue interface{}) []interface{} {
	s := reflect.ValueOf(mapValue)
	if s.Kind() != reflect.Map {
		panic("InterfaceMap: argument is not a map")
	}

	if s.IsNil() {
		return nil
	}

	keys := make([]interface{}, s.Len())
	for i, k := range s.MapKeys() {
		keys[i] = k.Interface()
	}
	return keys
}

// InterfaceMap is a map of interfaces.
func InterfaceMap(mapValue interface{}) map[interface{}]interface{} {
	s := reflect.ValueOf(mapValue)
	if s.Kind() != reflect.Map {
		panic("InterfaceMap: argument is not a map")
	}

	if s.IsNil() {
		return nil
	}

	keys := MapKeys(mapValue)
	ret := make(map[interface{}]interface{})
	for _, k := range keys {
		ret[k] = s.MapIndex(reflect.ValueOf(k)).Interface()
	}

	return ret
}
