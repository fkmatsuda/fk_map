package fkmap

import "testing"

func TestMapKeys(t *testing.T) {
	m := map[string]string{
		"a": "b",
		"c": "d",
	}
	keys := MapKeys(m)
	if len(keys) != 2 {
		t.Errorf("MapKeys: expected 2 keys, got %d", len(keys))
	}
	if keys[0] != "a" || keys[1] != "c" {
		t.Errorf("MapKeys: expected keys to be 'a' and 'c', got %v", keys)
	}
}

func TestMapKeys_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MapKeys: expected panic, got nil")
		}
	}()
	MapKeys("panic")
}

func TestMapKeys_nil(t *testing.T) {
	var m map[string]string
	keys := MapKeys(m)
	if keys != nil {
		t.Errorf("MapKeys: expected nil, got %v", keys)
	}
}

func TestInterfaceMap(t *testing.T) {
	m := map[string]string{
		"a": "b",
		"c": "d",
	}
	im := InterfaceMap(m)
	if len(im) != 2 {
		t.Errorf("InterfaceMap: expected 2 keys, got %d", len(im))
	}
	if im["a"] != "b" || im["c"] != "d" {
		t.Errorf("InterfaceMap: expected keys to be 'a' and 'c', got %v", im)
	}
}

func TestInterfaceMap_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("InterfaceMap: expected panic, got nil")
		}
	}()
	InterfaceMap("panic")
}

func TestInterfaceMap_nil(t *testing.T) {
	var m map[string]string
	im := InterfaceMap(m)
	if im != nil {
		t.Errorf("InterfaceMap: expected nil, got %v", im)
	}
}
