package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func (engine *engine) setFakeValue(key string) {
	engine.storage[key] = &fakeValueStore{}
}

type engineTestSuite struct {
	suite.Suite
	engine engine
}

func (suite *engineTestSuite) SetupTest() {
	suite.engine = *newEngine()
	// TODO: turn all of these strings into constants
	suite.engine.Set("intvalue", "1234567890")
	suite.engine.Set("stringvalue", "foobartestvalueaoeuhtns")
	suite.engine.setFakeValue("fakevalue")
}

func (suite *engineTestSuite) TestNewEngine() {
	var interfaceEngine Engine
	interfaceEngine = NewEngine()
	suite.NotNil(interfaceEngine)

	var rawEngine engine
	rawEngine = *newEngine()
	suite.NotNil(rawEngine)
}

func (suite *engineTestSuite) TestGetSetStringValues() {
	value, err := suite.engine.Get("stringvalue")
	suite.NoError(err)
	suite.Equal("foobartestvalueaoeuhtns", *value)

	err = suite.engine.Set("key", "abc")
	suite.NoError(err)

	value, err = suite.engine.Get("key")
	suite.NoError(err)
	suite.Equal("abc", *value)
}

func (suite *engineTestSuite) TestGetSetIntValues() {
	value, err := suite.engine.Get("intvalue")
	suite.NoError(err)
	suite.Equal("1234567890", *value)

	err = suite.engine.Set("key", "123")
	suite.NoError(err)

	value, err = suite.engine.Get("key")
	suite.NoError(err)
	suite.Equal("123", *value)
}

func (suite *engineTestSuite) TestGetOtherValues() {
	value, err := suite.engine.Get("doesnotexist")
	suite.NoError(err)
	suite.Nil(value)

	_, err = suite.engine.Get("fakevalue")
	suite.Error(err)
}

func (suite *engineTestSuite) TestGetSetMethod() {
	value, err := suite.engine.GetSet("intvalue", "9876")
	suite.NoError(err)
	suite.Equal("1234567890", *value)

	value, err = suite.engine.Get("intvalue")
	suite.NoError(err)
	suite.Equal("9876", *value)

	value, err = suite.engine.GetSet("stringvalue", "abc123")
	suite.NoError(err)
	suite.Equal("foobartestvalueaoeuhtns", *value)

	value, err = suite.engine.Get("stringvalue")
	suite.NoError(err)
	suite.Equal("abc123", *value)

	value, err = suite.engine.GetSet("doesnotexist", "it does now")
	suite.NoError(err)
	suite.Nil(value)

	value, err = suite.engine.Get("doesnotexist")
	suite.NoError(err)
	suite.Equal("it does now", *value)

	_, err = suite.engine.GetSet("fakevalue", "fake fake fake")
	suite.Error(err)

	_, err = suite.engine.Get("fakevalue")
	suite.Error(err)

	typeString, err := suite.engine.Type("fakevalue")
	suite.NoError(err)
	suite.Equal("fake", typeString)
}

func (suite *engineTestSuite) TestDel() {
	deleted, err := suite.engine.Del([]string{"stringvalue", "intvalue", "doesnotexist"})
	suite.NoError(err)
	suite.Equal(int64(2), deleted)

	value, err := suite.engine.Get("stringvalue")
	suite.NoError(err)
	suite.Nil(value)

	value, err = suite.engine.Get("intvalue")
	suite.NoError(err)
	suite.Nil(value)

	value, err = suite.engine.Get("doesnotexist")
	suite.NoError(err)
	suite.Nil(value)
}

func (suite *engineTestSuite) TestExists() {
	exists, err := suite.engine.Exists([]string{"stringvalue", "intvalue", "doesnotexist"})
	suite.NoError(err)
	suite.Equal(int64(2), exists)
}

func (suite *engineTestSuite) TestStrLen() {
	length, err := suite.engine.Strlen("stringvalue")
	suite.NoError(err)
	suite.Equal(int64(23), length)

	length, err = suite.engine.Strlen("intvalue")
	suite.NoError(err)
	suite.Equal(int64(10), length)

	length, err = suite.engine.Strlen("doesnotexist")
	suite.NoError(err)
	suite.Equal(int64(0), length)

	_, err = suite.engine.Strlen("fakevalue")
	suite.Error(err)
}

func (suite *engineTestSuite) TestIntIncrDecr() {
	err := suite.engine.Set("intvalue", "1")
	suite.NoError(err)

	value, err := suite.engine.Incr("intvalue")
	suite.NoError(err)
	suite.Equal(int64(2), value)

	// TODO: make Incrby and Decrby use int64 instead of string? :|
	value, err = suite.engine.Incrby("intvalue", "10")
	suite.NoError(err)
	suite.Equal(int64(12), value)

	value, err = suite.engine.Decr("intvalue")
	suite.NoError(err)
	suite.Equal(int64(11), value)

	value, err = suite.engine.Decrby("intvalue", "100")
	suite.NoError(err)
	suite.Equal(int64(-89), value)

	value, err = suite.engine.Incr("incr")
	suite.NoError(err)
	suite.Equal(int64(1), value)

	value, err = suite.engine.Incrby("incrby", "10")
	suite.NoError(err)
	suite.Equal(int64(10), value)

	value, err = suite.engine.Decr("decr")
	suite.NoError(err)
	suite.Equal(int64(-1), value)

	value, err = suite.engine.Decrby("decrby", "10")
	suite.NoError(err)
	suite.Equal(int64(-10), value)
}

func (suite *engineTestSuite) TestOtherIncrDecr() {
	_, err := suite.engine.Incr("stringvalue")
	suite.Error(err)

	_, err = suite.engine.Incrby("stringvalue", "10")
	suite.Error(err)

	_, err = suite.engine.Decr("stringvalue")
	suite.Error(err)

	_, err = suite.engine.Decrby("stringvalue", "10")
	suite.Error(err)

	// make sure we haven't modified the string value anywhere in there
	value, err := suite.engine.Get("stringvalue")
	suite.NoError(err)
	suite.Equal("foobartestvalueaoeuhtns", *value)

	_, err = suite.engine.Incrby("intvalue", "invalidint")
	suite.Error(err)

	_, err = suite.engine.Decrby("intvalue", "anotherinvalidint")
	suite.Error(err)

	// make sure we haven't modified the int value anywhere in there
	value, err = suite.engine.Get("intvalue")
	suite.NoError(err)
	suite.Equal("1234567890", *value)

	_, err = suite.engine.Incr("fakevalue")
	suite.Error(err)

	// TODO: there's some subtlety here in whether the string parsing fails or
	// the key is of the wrong type (string, fakevalue, etc) that could use some assertions
	// but right now we're not sending specific error types so we'll ignore it
	_, err = suite.engine.Incrby("fakevalue", "10")
	suite.Error(err)

	_, err = suite.engine.Decr("fakevalue")
	suite.Error(err)

	_, err = suite.engine.Decrby("fakevalue", "10")
	suite.Error(err)
}

func (suite *engineTestSuite) TestMGet() {

}

func (suite *engineTestSuite) TestMSet() {

}

func (suite *engineTestSuite) TestType() {

}

func (suite *engineTestSuite) TestInternalMethods() {
	// getOrDefault
	// getStore
	// set
	// del
}

func TestEngineTestSuite(t *testing.T) {
	suite.Run(t, new(engineTestSuite))
}

func TestMGet(t *testing.T) {
	engine := newEngine()

	stringValue := "aoeuhtns"
	intValue := "123"
	engine.Set("stringvalue", stringValue)
	engine.Set("intvalue", intValue)
	engine.setFakeValue("fake")

	values, err := engine.MGet([]string{"stringvalue", "intvalue", "doesnotexist"})
	assert.NoError(t, err)
	assert.Equal(t, []*string{&stringValue, &intValue, nil}, values)

	values, err = engine.MGet([]string{"doesnotexist", "doesnotexist2", "doesnotexist3"})
	assert.Equal(t, []*string{nil, nil, nil}, values)
	assert.NoError(t, err)

	// TODO: test what happens when mget with an invalid data type (e.g. fakeValueStore)
	values, err = engine.MGet([]string{"stringvalue", "doesnotexist", "fake"})
	assert.Equal(t, []*string{&stringValue, nil, nil}, values)
	assert.NoError(t, err)
}

func TestMSet(t *testing.T) {
	engine := newEngine()

	kvs := map[string]string{
		"foo": "aoeuhtns",
		"bar": "123",
	}

	err := engine.MSet(kvs)
	assert.NoError(t, err)

	value, err := engine.Get("foo")
	assert.NoError(t, err)
	assert.Equal(t, "aoeuhtns", *value)

	value, err = engine.Get("bar")
	assert.NoError(t, err)
	assert.Equal(t, "123", *value)
}

func TestType(t *testing.T) {
	engine := newEngine()
	engine.Set("stringvalue", "abc")

	typeString, err := engine.Type("stringvalue")
	assert.Equal(t, "string", typeString)
	assert.NoError(t, err)

	engine.Set("intvalue", "123")
	typeString, err = engine.Type("intvalue")
	assert.Equal(t, "string", typeString)
	assert.NoError(t, err)

	typeString, err = engine.Type("doesnotexist")
	assert.Equal(t, "none", typeString)
	assert.NoError(t, err)
}
