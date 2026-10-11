package main

import (
	"reflect"
	"testing"
)

// func Teste_EmptyStore(t *testing.T) {
// 	store := newStore()
// 	if len(store.data) != 0 {
// 		t.Errorf("Expected empty store, got %v", store.data)
// 	}
// }

func TestGet_limitedStore(t *testing.T) {
	store := newStore(2)
	store.Set("KEY1", "hihry")
	store.Set("KEY2", "hihry2")
	// Attempt to add a third key should fail
	_, err := store.Get("KEY3")	
	if err == nil {
		t.Errorf("Expected error when getting non-existent key, got nil")
	}
}	
func TestKeys(t *testing.T) {
	store := newStore()
	store.Set("KEY1", "hihry")
	store.Set("KEY2", "hihry2")
	store.Set("KEY3", "hihry3") // Overwrite existing
	keys := store.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}
	expectedKeys := map[string]bool{"KEY1": true, "KEY2": true, "KEY3": true}
	if !reflect.DeepEqual(keys, []string{"KEY1", "KEY2", "KEY3"}) {
		t.Errorf("Expected keys %v, got %v", expectedKeys, keys)
	}
}

func TestKeys_EmptyStore(t *testing.T) {
	store := newStore()
	keys := store.Keys()
	
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys, got %d", len(keys))
	}
	store.Set("KEY1", "value1")
}

// TODO : Assert Get("Hello") return ("world" , true) and Get("NonExistentKey") return ("", false)
func TestSetGet_RoundTrip(t *testing.T) {
	store := newStore()
	store.Set("hello", "world")
	if val, err := store.Get("hello"); err != nil || val != "world" {
		t.Errorf("Get() failed")
	}
	if val, err := store.Get("nonexistent"); err == nil || val != "" {
		t.Errorf("Expected Get('nonexistent') to return ('', false), got ('%s', %v)", val, err)
	}
}
