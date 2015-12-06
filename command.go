import main 

type Store map[string]interface{}

func (s *Store) StringSet(key string, value string) {
	s[key] = value
}

func (s *Store) StringGet(key string) interface {
	return s[key]
}

func (s *Store) ListPush(key string, value []string) {
	if len(s[key]) == 0 {
		s[key] = value
	} else {
		s[key] = append(s[key].([]string), value)
	} 
}

func (s *Store) ListRightPop(key string) interface {
	if len(s[key]) == 0 {
		return interface{}
	} else {
		v := s[key]
		return v[0]
	}
}


func (s *Store) ListRange(key string, start int, stop int) map[string]interface {
	if (start < 0 || stop < 0) {
		return 
	}


}