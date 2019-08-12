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

func (s *intValueStore) incrby(by int64) (int64, error) {
	s.intValue += by
	return s.intValue, nil
}

func (s *intValueStore) getType() string {
	// yes, it's an intValueStore, but redis returns "string" for this
	// I have a feeling that internally it's a string value always, but I dunno.
	return "string"
}
