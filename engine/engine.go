package engine

import "fmt"

type Engine struct {
	storage map[string]valueStore
}

func NewEngine() *Engine {
	e := &Engine{storage: make(map[string]valueStore)}
	return e
}

// TODO: probably turn this into an interface maybe? So we can have generic types?
// but at the same time, what interface do we want it to have? get/set doesn't make sense for everything
type valueStore struct {
	stringValue *string
	intValue    *int
}

func (engine *Engine) Get(key string) (*string, error) {
	store, ok := engine.storage[key]
	if !ok {
		return nil, nil
	}
	return store.get()
}

func (v *valueStore) get() (*string, error) {
	// if value := v.stringValue {
	// 	return value, nil
	// } elsif value := v.intValue {
	// 	return &fmt.Sprintf("%i", value), nil
	// }
	if v.stringValue != nil {
		return v.stringValue, nil
	} else if v.intValue != nil {
		value := fmt.Sprintf("%d", v.stringValue)
		return &value, nil
	}
	return nil, fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value") // thanks, redis ;-)
}

func (engine *Engine) Set(key string, value string) error {
	store, ok := engine.storage[key]
	if !ok {
		store = valueStore{}
		engine.storage[key] = store
	}
	store.set(value)
	return nil
}

func (v *valueStore) set(value string) error {
	v.stringValue = &value
	v.intValue = nil
	return nil
}
