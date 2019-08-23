package engine

import (
	"time"
)

type valueStore struct {
	expiresAt *time.Time
}

type valueStoreInterface interface {
	getType() string // to be implemented by "subclasses"
	expire(at *time.Time)
	expires() *time.Time
	expired() bool
}

type stringishValueStoreInterface interface {
	get() string
	incrby(by int64) (int64, error)
	len() int64
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
