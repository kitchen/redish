package engine

import "fmt"

type stringValueStore struct {
	stringValue string
	valueStore
}

func (s *stringValueStore) get() (string, error) {
	return s.stringValue, nil
}

func (s *stringValueStore) incrby(by int64) (int64, error) {
	return 0, fmt.Errorf("ERR value is not an integer or out of range")
}

func (s *stringValueStore) getType() string {
	return "string"
}
