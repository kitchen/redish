package engine

import "fmt"

type stringValueStore struct {
	stringValue string
	valueStore
}

func (s *stringValueStore) get() string {
	return s.stringValue
}

func (s *stringValueStore) incrby(by int64) (int64, error) {
	return 0, fmt.Errorf("ERR value is not an integer or out of range")
}

func (s *stringValueStore) len() int64 {
	return int64(len(s.stringValue))
}

func (s *stringValueStore) getType() string {
	return "string"
}
