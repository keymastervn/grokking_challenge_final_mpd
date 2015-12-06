package main

func (s *Store) SADD(key string, values ...string) string {
	var set = make(map[string]interface{})
	var dvalues []string
	for _, v := range values {
		if _, ok := set[v]; ok {
			continue
		}

		dvalues = append(dvalues, v)
	}

	if _, ok := (*s)[key]; ok {
		values = (*s)[key].([]string)
		for _, v := range values {
			if _, ok := set[v]; ok {
				continue
			}

			dvalues = append(dvalues, v)
		}
	}

	(*s)[key] = dvalues
	return OK
}

func (s *Store) SCARD(key string) int {
	if _, ok := (*s)[key]; ok {
		return len((*s)[key].([]string))
	}
	return 0
}

func (s *Store) SMEMBERS(key string) []string {
	if _, ok := (*s)[key]; ok {
		return (*s)[key].([]string)
	}
	return []string{}
}

func (s *Store) SREM(key string, values ...string) string {
	var dvalues []string
	if _, ok := (*s)[key]; ok {
		storeValues := (*s)[key].([]string)
		for _, sv := range storeValues {
			for _, v := range values {
				if sv != v {
					dvalues = append(dvalues, v)
				}
			}
		}

		(*s)[key] = dvalues
	}

	return OK
}

func (s *Store) SINTER(keys ...string) []string {
	var result []string
	var resultMap = make(map[string]int)
	for _, key := range keys {
		if _, ok := (*s)[key]; !ok {
			return []string{}
		}

		values := (*s)[key].([]string)
		for _, v := range values {
			resultMap[v] = resultMap[v] + 1
		}
	}

	l := len(keys)
	for k, v := range resultMap {
		if v != l {
			continue
		}

		result = append(result, k)
	}

	return result
}
