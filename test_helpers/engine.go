package test_helpers

const stringValue = "foobartestvaluehtnsaoeu"
const stringValueKey = "stringvaluekey"
const intValueString = "1234567890"
const intValueKey = "intvaluekey"
const fakeValueKey = "fakevaluekey"
const doesNotExistKey = "doesnotexistkey"

func (engine *engine) setFakeValue(key string) {
	engine.storage[key] = &fakeValueStore{}
}
