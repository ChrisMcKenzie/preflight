package preflight

import "sync"

type Set struct {
	m    map[interface{}]interface{}
	once sync.Once
}

type Hashable interface {
	Hashcode() interface{}
}

func hashcode(v interface{}) interface{} {
	if h, ok := v.(Hashable); ok {
		return h.Hashcode()
	}

	return v
}

func (s *Set) Add(v interface{}) {
	s.once.Do(s.init)
	s.m[hashcode(v)] = v
}

func (s *Set) Delete(v interface{}) {
	s.once.Do(s.init)
	delete(s.m, hashcode(v))
}

func (s *Set) Include(v interface{}) bool {
	s.once.Do(s.init)
	_, ok := s.m[hashcode(v)]
	return ok
}

func (s *Set) List() []interface{} {
	if s == nil {
		return nil
	}

	r := make([]interface{}, 0, len(s.m))
	for _, v := range s.m {
		r = append(r, v)
	}

	return r
}

func (s *Set) Len() int {
	if s == nil {
		return 0
	}

	return len(s.m)
}

func (s *Set) Intersection(o *Set) *Set {
	result := new(Set)
	if s == nil {
		return result
	}

	if o != nil {
		for _, v := range s.m {
			if o.Include(v) {
				result.Add(v)
			}
		}
	}

	return result
}

func (s *Set) init() {
	s.m = make(map[interface{}]interface{})
}
