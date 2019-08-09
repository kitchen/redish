package engine

import (
	"fmt"
	"log"
	"strconv"
)

type Engine struct {
	storage map[string]*valueStore
}

func NewEngine() *Engine {
	e := &Engine{storage: make(map[string]*valueStore)}
	return e
}

// TODO: probably turn this into an interface maybe? So we can have generic types?
// but at the same time, what interface do we want it to have? get/set doesn't make sense for everything
type valueStore struct {
	stringValue *string
	intValue    *int64
}

func (engine *Engine) Get(key string) (*string, error) {
	store, ok := engine.storage[key]
	if !ok {
		log.Printf("key %v doesn't exist!", key)
		return nil, nil
	}
	log.Printf("key %v exists, got store: %v", key, store)
	return store.get()
}

func (v *valueStore) get() (*string, error) {
	// if value := v.stringValue {
	// 	return value, nil
	// } elsif value := v.intValue {
	// 	return &fmt.Sprintf("%i", value), nil
	// }
	log.Printf("value: %v", v)
	if v.stringValue != nil {
		return v.stringValue, nil
	} else if v.intValue != nil {
		value := fmt.Sprintf("%d", *v.intValue)
		return &value, nil
	}
	return nil, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value") // thanks, redis ;-)
}

func (engine *Engine) Set(key string, value string) error {
	store, ok := engine.storage[key]
	if !ok {
		log.Printf("key %v doesn't exist yet, let's make it", key)
		store = &valueStore{}
		engine.storage[key] = store
	}
	store.set(value)
	store = engine.storage[key]
	log.Printf("store after setting: %v", store)
	return nil
}

func (v *valueStore) set(value string) error {
	log.Printf("ourselves before: %v", v)
	if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
		log.Printf("value being set looks like an int: %v -> %v", value, intValue)
		v.intValue = &intValue
		v.stringValue = nil
	} else {
		log.Printf("value being set doesn't look like an int: %v", value)
		v.intValue = nil
		v.stringValue = &value
	}

	log.Printf("stuff!")
	log.Printf("ourselves! %v", v)
	return nil
}
