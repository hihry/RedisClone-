package main

import (
	"errors"
	"sort"
)

var (
	ErrorType  = errors.New("key not found")
	ErrStoreFull = errors.New("store is full")
)

type Store struct {
	data    map[string]string
	maxKeys int
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Sort the keys for consistent order
	return keys
}

func (s *Store) Get(key string) (string, error) {
	val, ok := s.data[key]
	if !ok {
		return "", ErrorType
	}
	return val, nil
}
func (s *Store) Set(key string, val string) (string, bool) {
	if _, exists := s.data[key]; !exists &&
		s.maxKeys > 0 && len(s.data) >= s.maxKeys {
		return "", false
	}
	s.data[key] = val
	return val, true
}

func newStore(maxKeys ...int) *Store {
	limit := 0
	if len(maxKeys) > 0 {
		limit = maxKeys[0]
	}

	return &Store{
		data:    make(map[string]string),
		maxKeys: limit,
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
