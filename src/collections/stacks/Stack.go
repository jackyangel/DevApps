package stacks

import (
	"errors"
	"sync"
)

type Stack2 struct {
	lock  sync.Mutex //you don't have to do this if you don't want thread safety
	s     []string
	Count int
}

func NewStack() *Stack2 {
	return &Stack2{sync.Mutex{}, make([]string, 0), 0}
}

func (s *Stack2) Push(v string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
	s.Count++
}

func (s *Stack2) Pop() (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return "", errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	s.Count--
	return res, nil
}
