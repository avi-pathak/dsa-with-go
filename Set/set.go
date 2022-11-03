package set

import (
	"golang.org/x/exp/constraints"
)

type Set[T constraints.Ordered] struct {
	set map[T]T
}

func (s *Set[t]) Add(value t) {
	if s.set == nil {
		s.set = make(map[t]t)
	}
	s.set[value] = value
}

func (s *Set[t]) Delete(key t) {
	delete(s.set, key)
}

func (s *Set[t]) Has(key t) bool {
	_, ok := s.set[key]

	return ok

}
func (s *Set[t]) Count() int {
	return len(s.set)
}
