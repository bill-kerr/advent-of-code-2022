package util

var exists = struct{}{}

type Set struct {
	m map[int]struct{}
}

func NewSet() *Set {
	s := &Set{}
	s.m = make(map[int]struct{})
	return s
}

func (s *Set) Add(value int) {
	s.m[value] = exists
}

func (s *Set) Remove(value int) {
	delete(s.m, value)
}

func (s *Set) Has(value int) bool {
	_, ok := s.m[value]
	return ok
}

func (s *Set) ToMap() map[int]struct{} {
	return s.m
}