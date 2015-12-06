package main

import "errors"

const OK = "OK"

func (s *Store) LLEN(key string) int {
	return len((*s)[key].([]interface{}))
}

func (s *Store) RPUSH(key string, value ...string) string {
	if _, ok := (*s)[key]; !ok {
		(*s)[key] = value
	} else {
		(*s)[key] = append((*s)[key].([]string), value...)
	}
	return OK
}

func (s *Store) LPOP(key string) string {
	if _, ok := (*s)[key]; !ok {
		return ""
	}
	a := (*s)[key].([]string)
	(*s)[key] = a[1:]

	return a[0]
}

func (s *Store) RPOP(key string) string {
	if _, ok := (*s)[key]; !ok {
		return ""
	}

	a := (*s)[key].([]string)
	if len(a) > 1 {
		(*s)[key] = a[:len(a)-1]
	} else {
		delete(*s, key)
	}
	return a[len(a)-1]
}

func (s *Store) LRANGE(key string, start int, stop int) ([]string, error) {
	if start < 0 || stop < 0 {
		return nil, errors.New("EINV")
	}

	if _, ok := (*s)[key]; !ok {
		return []string{}, nil
	}

	a := (*s)[key].([]string)
	if len(a) > stop {
		return a[start:stop], nil
	}

	return nil, errors.New("EINV")
}
