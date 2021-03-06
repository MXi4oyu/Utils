package Set

import (
	"sync"
)

type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func New() *Set {

	return &Set{
		m: map[interface{}]bool{},
	}
}

func (s *Set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Remove(item interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *Set) Has(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
    if _, ok := s.m[item];ok{

		return true
	}else{
		return false
	}
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]bool{}
}

func (s *Set) List() []interface{} {

	s.RLock()
	defer s.RUnlock()
	list := []interface{}{}
	for item := range s.m {
		list = append(list, item)
	}

	return list
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) IsEmpty() bool {

	if s.Len() == 0 {
		return true
	} else {
		return false
	}
}
