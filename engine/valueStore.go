package engine

import "fmt"

type valueStore struct{}
type valueStoreInterface interface {
	get() (string, error)
	incrby(by int64) (string, error)
}

func (s *valueStore) get() (string, error) {
	return "", fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value") // thanks, redis ;-)
}

func (s *valueStore) incrby(by int64) (string, error) {
	return "", fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
}
