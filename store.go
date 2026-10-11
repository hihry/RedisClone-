package main

import "sort"

type Store struct {
	data map[string]string
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Sort the keys for consistent order
	return keys
}

func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}
func (s *Store) Set(key string, val string) (string, bool) {
	s.data[key] = val
	return val, true
}
func newStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Delete(key string) bool {
	_, ok := s.data[key]
	if ok {
		delete(s.data, key)
		return true
	}
	return false
}
