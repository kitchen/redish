package engine

import (
	"fmt"
	"strconv"
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
	store, _ := engine.storage[key]
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
	store := engine.getStore(key)
	if store != nil {
		return store
	}
	return engine.set(key, defaultValue)
}

// TODO: should probably lock the db
func (engine *engine) Del(keys []string) (int64, error) {
	deleted := int64(0)
	for _, key := range keys {
		if _, ok := engine.storage[key]; ok {
			delete(engine.storage, key)
			deleted += 1
		}
	}
	return deleted, nil
}

func (engine *engine) Exists(keys []string) (int64, error) {
	exists := int64(0)
	for _, key := range keys {
		if _, ok := engine.storage[key]; ok {
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
