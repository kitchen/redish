package engine

import (
	"log"
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
}

func NewEngine() Engine {
	e := &engine{storage: make(map[string]valueStoreInterface)}
	return e
}

func (engine *engine) Get(key string) (*string, error) {
	store, ok := engine.storage[key]
	if !ok {
		log.Printf("key %v doesn't exist!", key)
		return nil, nil
	}
	log.Printf("key %v exists, got store: %v", key, store)
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
