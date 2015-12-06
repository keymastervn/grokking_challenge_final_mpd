package main

type Store map[string]interface{}

func (s *Store) StringSet(key string, value string) string {
	(*s)[key] = value
	return OK
}

func (s *Store) StringGet(key string) string {
	return (*s)[key].(string)
}
