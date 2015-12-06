package main

type Store map[string]interface{}

func (s *Store) StringSet(key string, value string) {
	s[key] = value
}

func (s *Store) StringGet(key string) interface {
	return s[key]
}


