package engine

import (
	"fmt"
)

type intValueStore struct {
	intValue int64
	valueStore
}

func (s *intValueStore) get() (string, error) {
	return fmt.Sprintf("%d", s.intValue), nil
}

func (s *intValueStore) incrby(by int64) (string, error) {
	s.intValue += by
	return s.get()
}
