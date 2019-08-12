package engine

import "fmt"

type valueStore struct{}
type valueStoreInterface interface {
	get() (string, error)
	incrby(by int64) (int64, error)
	getType() string // to be implemented by "subclasses"
}

func (s *valueStore) get() (string, error) {
	return "", fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value") // thanks, redis ;-)
}

func (s *valueStore) incrby(by int64) (int64, error) {
	return 0, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
}
