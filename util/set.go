package util

var exists = struct{}{}

type SetValue interface {
	comparable
}

type Set[V SetValue] struct {
	m map[V]struct{}
}

func NewSet[V SetValue]() *Set[V] {
	s := &Set[V]{}
	s.m = make(map[V]struct{})
	return s
}

func (s *Set[V]) Add(value V) {
	s.m[value] = exists
}

func (s *Set[V]) Remove(value V) {
	delete(s.m, value)
}

func (s *Set[V]) Has(value V) bool {
	_, ok := s.m[value]
	return ok
}

func (s *Set[V]) ToMap() map[V]struct{} {
	return s.m
}

func (s *Set[V]) Size() int {
	return len(s.m)
}