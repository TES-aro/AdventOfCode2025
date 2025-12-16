package utilsAOC

import (

)

type Set [T comparable] struct{
	innerMap map[T]bool
}

func (s *Set[T]) Add(item T){
	s.innerMap[item] = true
}

func (s *Set[T]) Remove(item T){
	delete(s.innerMap, item)
}

func (s *Set[T]) Contains(item T) bool {
	if s.innerMap[item] == true{
		return true
	}
	return false
}

func NewSet[T comparable]() *Set[T] {
	newSet := new(Set[T])
	newSet.innerMap = map[T]bool{}
	return newSet
}

func (s *Set[T]) Values() []T{
	i := 0
	keys := make([]T, len(s.innerMap))
	for k := range s.innerMap{
		keys[i] = k
		i++
	}
	return keys
}

