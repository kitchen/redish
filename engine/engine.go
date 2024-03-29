package engine

import (
	"fmt"
	"strconv"
	"time"
)

type engine struct {
	storage map[string]valueStoreInterface
}

type Engine interface {
	Get(key string) (*string, error)
	Set(key string, value string) error
	Del(keys []string) (int64, error)
	Exists(keys []string) (int64, error)
	Incr(key string) (int64, error)
	Decr(key string) (int64, error)
	Incrby(key string, by int64) (int64, error)
	Decrby(key string, by int64) (int64, error)
	Strlen(key string) (int64, error)
	GetSet(key string, value string) (*string, error)
	MGet(keys []string) ([]*string, error)
	MSet(kvs map[string]string) error
	Type(key string) (string, error)
	Expire(key string, seconds int64) (bool, error)
	PExpire(key string, millis int64) (bool, error)
	// seconds since epoch
	ExpireAt(key string, seconds int64) (bool, error)
	// millis since epoch
	PExpireAt(key string, millis int64) (bool, error)
	Persist(key string) (bool, error)
	TTL(key string) (int64, error)
	PTTL(key string) (int64, error)
}

func NewEngine() Engine {
	return newEngine()
}

func newEngine() *engine {
	e := &engine{storage: make(map[string]valueStoreInterface)}
	return e
}

func (engine *engine) Get(key string) (*string, error) {
	store := engine.getStore(key)
	if store == nil {
		return nil, nil
	}

	if store, ok := store.(stringishValueStoreInterface); ok {
		value := store.get()
		return &value, nil
	}
	return nil, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
}

func (engine *engine) getStore(key string) valueStoreInterface {
	store, _ := engine.storage[key]
	if store == nil {
		return nil
	}

	if store.expired() {
		engine.del(key)
		return nil
	}
	return store
}

func (engine *engine) Set(key string, value string) error {
	engine.set(key, value)
	return nil
}

func (engine *engine) set(key string, value string) valueStoreInterface {
	var storage valueStoreInterface
	if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
		storage = &intValueStore{intValue: intValue}
	} else {
		storage = &stringValueStore{stringValue: value}
	}
	engine.storage[key] = storage
	return storage
}

func (engine *engine) getOrDefault(key string, defaultValue string) valueStoreInterface {
	if store := engine.getStore(key); store != nil {
		return store
	}
	return engine.set(key, defaultValue)
}

// TODO: should probably lock the db
func (engine *engine) Del(keys []string) (int64, error) {
	deleted := int64(0)
	for _, key := range keys {
		if engine.del(key) {
			deleted++
		}
	}
	return deleted, nil
}

func (engine *engine) del(key string) bool {
	if _, ok := engine.storage[key]; ok {
		delete(engine.storage, key)
		return true
	}
	return false
}

func (engine *engine) Exists(keys []string) (int64, error) {
	exists := int64(0)
	for _, key := range keys {
		if store := engine.getStore(key); store != nil {
			exists++
		}
	}
	return exists, nil
}

func (engine *engine) Incr(key string) (int64, error) {
	store := engine.getOrDefault(key, "0")

	if store, ok := store.(stringishValueStoreInterface); ok {
		return store.incrby(1)
	}
	return 0, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
}

func (engine *engine) Decr(key string) (int64, error) {
	store := engine.getOrDefault(key, "0")
	if store, ok := store.(stringishValueStoreInterface); ok {
		return store.incrby(-1)
	}
	return 0, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
}

func (engine *engine) Incrby(key string, by int64) (int64, error) {
	store := engine.getOrDefault(key, "0")
	if store, ok := store.(stringishValueStoreInterface); ok {
		return store.incrby(by)
	}
	return 0, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
}

func (engine *engine) Decrby(key string, by int64) (int64, error) {
	store := engine.getOrDefault(key, "0")
	if store, ok := store.(stringishValueStoreInterface); ok {
		return store.incrby(-by)
	}
	return 0, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
}

func (engine *engine) Strlen(key string) (int64, error) {
	store := engine.getStore(key)
	if store == nil {
		return 0, nil
	}

	if castStore, ok := store.(stringishValueStoreInterface); ok {
		return castStore.len(), nil
	}
	return 0, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
}

func (engine *engine) GetSet(key string, newValue string) (*string, error) {
	store := engine.getStore(key)
	var oldValue *string
	if store == nil {
		oldValue = nil
	} else {
		if store, ok := store.(stringishValueStoreInterface); ok {
			tmpValue := store.get()
			oldValue = &tmpValue
		} else {
			return nil, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
		}
	}
	engine.set(key, newValue)
	return oldValue, nil
}

func (engine *engine) MGet(keys []string) ([]*string, error) {
	values := make([]*string, len(keys))
	for i, key := range keys {
		if store := engine.getStore(key); store != nil {
			if store, ok := store.(stringishValueStoreInterface); ok {
				value := store.get()
				values[i] = &value
			} else {
				values[i] = nil
			}
		} else {
			values[i] = nil
		}
	}
	return values, nil
}

func (engine *engine) MSet(kvs map[string]string) error {
	for k, v := range kvs {
		engine.set(k, v)
	}

	return nil
}

func (engine *engine) Type(key string) (string, error) {
	if store := engine.getStore(key); store != nil {
		return store.getType(), nil
	}
	return "none", nil
}

func (engine *engine) Expire(key string, seconds int64) (bool, error) {
	duration := time.Duration(time.Second.Nanoseconds() * seconds)
	expireAt := time.Now().Add(duration)
	return engine.expireAt(key, expireAt), nil
}
func (engine *engine) PExpire(key string, millis int64) (bool, error) {
	duration := time.Duration(time.Millisecond.Nanoseconds() * millis)
	expireAt := time.Now().Add(duration)
	return engine.expireAt(key, expireAt), nil
}

func (engine *engine) ExpireAt(key string, seconds int64) (bool, error) {
	expireAt := time.Unix(seconds, 0)
	return engine.expireAt(key, expireAt), nil
}

func (engine *engine) PExpireAt(key string, millis int64) (bool, error) {
	millisPart := millis % 1000
	secondsPart := millis / 1000
	millisNanos := time.Millisecond.Nanoseconds() * millisPart
	expireAt := time.Unix(secondsPart, millisNanos)
	return engine.expireAt(key, expireAt), nil
}

func (engine *engine) expireAt(key string, at time.Time) bool {
	store := engine.getStore(key)
	if store == nil {
		return false
	}

	if time.Now().After(at) {
		engine.del(key)
		return false
	} else {
		store.expire(&at)
		return true
	}
}

func (engine *engine) Persist(key string) (bool, error) {
	store := engine.getStore(key)
	if store == nil || store.expires() == nil {
		return false, nil
	}
	store.expire(nil)
	return true, nil
}

func (engine *engine) TTL(key string) (int64, error) {
	store := engine.getStore(key)
	if store == nil {
		return int64(-2), nil
	}
	expireAt := store.expires()
	if expireAt == nil {
		return int64(-1), nil
	}

	ttl := expireAt.Sub(time.Now()).Seconds()
	return int64(ttl), nil
}

func (engine *engine) PTTL(key string) (int64, error) {
	store := engine.getStore(key)
	if store == nil {
		return int64(-2), nil
	}
	expireAt := store.expires()
	if expireAt == nil {
		return int64(-1), nil
	}

	ttl := expireAt.Sub(time.Now()).Nanoseconds() / 1000000
	return ttl, nil
}
