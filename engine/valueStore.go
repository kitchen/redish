package engine

import (
	"fmt"
	"time"
)

type valueStore struct {
	expiresAt *time.Time
}

type valueStoreInterface interface {
	get() (string, error)
	incrby(by int64) (int64, error)
	getType() string // to be implemented by "subclasses"
	expire(at *time.Time)
	expires() *time.Time
	expired() bool
}

func (s *valueStore) get() (string, error) {
	return "", fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value") // thanks, redis ;-)
}

func (s *valueStore) incrby(by int64) (int64, error) {
	return 0, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
}

func (s *valueStore) expire(at *time.Time) {
	s.expiresAt = at
}

func (s *valueStore) expires() *time.Time {
	return s.expiresAt
}

func (s *valueStore) expired() bool {
	if s.expiresAt == nil {
		return false
	}
	return time.Now().After(*s.expiresAt)
}
