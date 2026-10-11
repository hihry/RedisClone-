package main

import (
	"testing"
	"reflect"
)

func TestKeys(t *testing.T){
	store := newStore()
	store.Set("KEY1" , "hihry")
	store.Set("KEY2" , "hihry2")
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