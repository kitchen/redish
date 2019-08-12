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
	Incr(key string) (string, error)
	Decr(key string) (string, error)
	Incrby(key string, by string) (string, error)
	Decrby(key string, by string) (string, error)
	Strlen(key string) (string, error)
	GetSet(key string, value string) (*string, error)
	MGet(keys []string) ([]*string, error)
	MSet(kvs map[string]string) error
	Type(key string) string
	Expire(key string, seconds int64) bool
	Persist(key string) bool
	TTL(key string) int64
}

func NewEngine() Engine {
	e := &engine{storage: make(map[string]valueStoreInterface)}
	return e
}

func (engine *engine) Get(key string) (*string, error) {
	store := engine.getStore(key)
	if store == nil {
		return nil, nil
	}
	value, err := store.get()
	return &value, err
}

func (engine *engine) getStore(key string) valueStoreInterface {
	if store, _ := engine.storage[key]; store != nil {
		if store.expired() {
			engine.del(key)
			return nil
		}
		return store
	}
	return nil
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

// func (engine *engine) setex(key string, expiration *Date) err {
//
// }

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
			deleted += 1
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
			exists += 1
		}
	}
	return exists, nil
}

func (engine *engine) Incr(key string) (string, error) {
	store := engine.getOrDefault(key, "0")

	return store.incrby(1)
}

func (engine *engine) Decr(key string) (string, error) {
	store := engine.getOrDefault(key, "0")
	return store.incrby(-1)
}

func (engine *engine) Incrby(key string, by string) (string, error) {
	if intValue, err := strconv.ParseInt(by, 10, 64); err == nil {
		store := engine.getOrDefault(key, "0")
		return store.incrby(intValue)
	}
	return "", fmt.Errorf("ERR value is not an integer or out of range")

}

func (engine *engine) Decrby(key string, by string) (string, error) {
	if intValue, err := strconv.ParseInt(by, 10, 64); err == nil {
		store := engine.getOrDefault(key, "0")
		return store.incrby(-intValue)
	}
	return "", fmt.Errorf("ERR value is not an integer or out of range")
}

func (engine *engine) Strlen(key string) (string, error) {
	store := engine.getStore(key)
	if store == nil {
		return "0", nil
	}
	if value, err := store.get(); err == nil {
		return fmt.Sprintf("%d", len(value)), nil
	} else {
		return "", err
	}
}

func (engine *engine) GetSet(key string, value string) (*string, error) {
	store := engine.getStore(key)
	var oldValue *string
	if store == nil {
		oldValue = nil
	} else {
		value, err := store.get()
		if err != nil {
			return nil, err
		}
		oldValue = &value
	}
	engine.set(key, value)
	return oldValue, nil
}

func (engine *engine) MGet(keys []string) ([]*string, error) {
	values := make([]*string, len(keys))
	for i, key := range keys {
		if store := engine.getStore(key); store != nil {
			if value, err := store.get(); err == nil {
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

func (engine *engine) Type(key string) string {
	if store := engine.getStore(key); store != nil {
		return store.getType()
	}
	return "none"
}

func (engine *engine) Expire(key string, seconds int64) bool {
	if seconds <= 0 {
		return engine.del(key)
	}
	if store := engine.getStore(key); store != nil {
		expires := time.Now().Add(0)
		store.expire(&expires)
		return true
	}
	return false
}

func (engine *engine) Persist(key string) bool {
	store := engine.getStore(key)
	if store == nil || store.expires() == nil {
		return false
	}
	store.expire(nil)
	return true
}

func (engine *engine) TTL(key string) int64 {
	store := engine.getStore(key)
	if store == nil {
		return int64(-2)
	}
	if expires := store.expires(); expires != nil {
		duration := expires.Sub(time.Now())
		return int64(duration.Truncate(time.Second).Seconds())
	} else {
		return int64(-1)
	}
}
