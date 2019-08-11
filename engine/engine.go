package engine

import (
	"fmt"
	"log"
	"strconv"
)

// TODO: convert this to an interface and make Engine struct be engine and unexported
// https://stackoverflow.com/questions/37135193/how-to-set-default-values-in-go-structs
// this forces users to use the NewEngine() method which sets default values and such
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

// TODO: should probably lock the db
func (engine *Engine) Del(keys []string) (*int64, error) {
	deleted := int64(0)
	for _, key := range keys {
		if _, ok := engine.storage[key]; ok {
			delete(engine.storage, key)
			deleted += 1
		}
	}
	return &deleted, nil
}

func (engine *Engine) Incr(key string) (*int64, error) {
	store, ok := engine.storage[key]
	if !ok {
		value := int64(0)
		store = &valueStore{intValue: &value}
		engine.storage[key] = store
	}
	return store.incr()
}

// probably DRY this out into a single incrby function and the others just specify their by
// and DRY it out at the Engine layer, too, then there's just one function on the valueStore
// and everything collapses to the Engine level, which is also the high level implementation of Incrby
// TODO: check for out of bounds condition, I'm assuming golang will either blow up sad like or overflow this if we try to out of bounds
// and I'd prefer to just error that back to the user
func (s *valueStore) incr() (*int64, error) {
	if s.intValue == nil {
		return nil, fmt.Errorf("ERR value is not an integer or out of range")
	}

	// this should probably lock the key briefly
	var intValue int64
	intValue = *s.intValue
	intValue += 1
	s.intValue = &intValue
	return s.intValue, nil
}

func (engine *Engine) Decr(key string) (*int64, error) {
	store, ok := engine.storage[key]
	if !ok {
		value := int64(0)
		store = &valueStore{intValue: &value}
		engine.storage[key] = store
	}
	return store.decr()
}
func (s *valueStore) decr() (*int64, error) {
	if s.intValue == nil {
		return nil, fmt.Errorf("ERR value is not an integer or out of range")
	}

	// this should probably lock the key briefly
	var intValue int64
	intValue = *s.intValue
	intValue -= 1
	s.intValue = &intValue
	return s.intValue, nil
}
