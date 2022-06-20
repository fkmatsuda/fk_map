package fkmap

import "reflect"

// MapKeys is a slice of keys of a map.
func MapKeys[K comparable, V any](mapValue map[K]V) []K {
	s := reflect.ValueOf(mapValue)

	if s.IsNil() {
		return nil
	}

	keys := make([]K, s.Len())
	for i, k := range s.MapKeys() {
		keys[i] = k.Interface().(K)
	}
	return keys
}

// InterfaceMap is a map of interfaces.
func InterfaceMap[K comparable, V any](mapValue map[K]V) map[interface{}]interface{} {
	s := reflect.ValueOf(mapValue)
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
