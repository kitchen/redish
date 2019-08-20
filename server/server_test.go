package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/kitchen/redish/redish"
	"github.com/stretchr/testify/suite"
)

var stringValue = "stringValue"
var stringValueKey = "stringvalue"
var stringValueKeyObject = redish.Key{Key: stringValueKey}
var stringKeyValue = redish.KeyValue{Key: stringValueKey, Value: stringValue}
var intValue = "424242"
var intValueKey = "intvalue"
var intValueKeyObject = redish.Key{Key: intValueKey}
var intKeyValue = redish.KeyValue{Key: intValueKey, Value: intValue}
var doesNotExistKey = "doesnotexist"
var doesNotExistKeyObject = redish.Key{Key: doesNotExistKey}
var fakeValueKey = "fakevalue"
var fakeValueKeyObject = redish.Key{Key: fakeValueKey}

type serverTestSuite struct {
	engine  *mockEngine
	server  *redishServer
	context context.Context
	suite.Suite
}

func (suite *serverTestSuite) SetupTest() {
	suite.engine = &mockEngine{}
	suite.server = newServer(suite.engine)
	ctx := context.TODO()
	suite.context = ctx
}

func (suite *serverTestSuite) TestDoesNotExistGet() {
	suite.engine.On("Get", doesNotExistKey).Return(nil, nil)
	value, err := suite.server.Get(suite.context, &doesNotExistKeyObject)
	suite.NotNil(value)
	suite.Nil(value.Value)
	suite.NoError(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestStringValueGet() {
	suite.engine.On("Get", stringValueKey).Return(&stringValue, nil)
	value, err := suite.server.Get(suite.context, &stringValueKeyObject)
	suite.NoError(err)
	suite.Equal(stringValue, value.Value.Value)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestIntValueGet() {
	suite.engine.On("Get", intValueKey).Return(&intValue, nil)
	value, err := suite.server.Get(suite.context, &intValueKeyObject)
	suite.NoError(err)
	suite.Equal(intValue, value.Value.Value)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestFakeValueGet() {
	suite.engine.On("Get", fakeValueKey).Return(nil, fmt.Errorf("invalid value"))
	_, err := suite.server.Get(suite.context, &fakeValueKeyObject)
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestSet() {
	suite.engine.On("Set", stringValueKey, stringValue).Return(nil)
	ret, err := suite.server.Set(suite.context, &redish.SetRequest{Key: stringValueKey, Value: stringValue})
	suite.Equal(redish.OK{}, *ret)
	suite.NoError(err)

	suite.engine.On("Set", intValueKey, intValue).Return(fmt.Errorf("somehow this failed"))
	_, err = suite.server.Set(suite.context, &redish.SetRequest{Key: intValueKey, Value: intValue})
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestDele() {
	suite.engine.On("Del", []string{stringValueKey, intValueKey}).Return(2, nil)
	ret, err := suite.server.Dele(suite.context, &redish.KeyList{Keys: []*redish.Key{&stringValueKeyObject, &intValueKeyObject}})
	suite.Equal(int64(2), ret.Value)
	suite.NoError(err)

	suite.engine.On("Del", []string{doesNotExistKey}).Return(2, fmt.Errorf("somehow this failed"))
	_, err = suite.server.Dele(suite.context, &redish.KeyList{Keys: []*redish.Key{&doesNotExistKeyObject}})
	suite.Error(err)

	suite.engine.On("Del", []string{}).Return(0, nil)
	ret, err = suite.server.Dele(suite.context, &redish.KeyList{Keys: []*redish.Key{}})
	suite.Equal(int64(0), ret.Value)
	suite.NoError(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestExists() {
	suite.engine.On("Exists", []string{stringValueKey, intValueKey, doesNotExistKey}).Return(2, nil)
	ret, err := suite.server.Exists(suite.context, &redish.KeyList{Keys: []*redish.Key{&stringValueKeyObject, &intValueKeyObject, &doesNotExistKeyObject}})
	suite.NoError(err)
	suite.Equal(int64(2), ret.Value)

	suite.engine.On("Exists", []string{fakeValueKey}).Return(0, fmt.Errorf("somehow this failed"))
	_, err = suite.server.Exists(suite.context, &redish.KeyList{Keys: []*redish.Key{&fakeValueKeyObject}})
	suite.Error(err)

	suite.engine.On("Exists", []string{}).Return(0, nil)
	ret, err = suite.server.Exists(suite.context, &redish.KeyList{Keys: []*redish.Key{}})
	suite.Equal(int64(0), ret.Value)
	suite.NoError(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestIncr() {
	suite.engine.On("Incr", intValueKey).Return(2, nil)
	ret, err := suite.server.Incr(suite.context, &intValueKeyObject)
	suite.NoError(err)
	suite.Equal(int64(2), ret.Value)

	suite.engine.On("Incr", stringValueKey).Return(0, fmt.Errorf("invalid data type"))
	_, err = suite.server.Incr(suite.context, &stringValueKeyObject)
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestDecr() {
	suite.engine.On("Decr", intValueKey).Return(2, nil)
	ret, err := suite.server.Decr(suite.context, &intValueKeyObject)
	suite.NoError(err)
	suite.Equal(int64(2), ret.Value)

	suite.engine.On("Decr", stringValueKey).Return(0, fmt.Errorf("invalid data type"))
	_, err = suite.server.Decr(suite.context, &stringValueKeyObject)
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestIncrby() {
	suite.engine.On("Incrby", intValueKey, int64(2)).Return(2, nil)
	ret, err := suite.server.Incrby(suite.context, &redish.KeyIntValue{Key: intValueKey, Value: int64(2)})
	suite.NoError(err)
	suite.Equal(int64(2), ret.Value)

	suite.engine.On("Incrby", stringValueKey, int64(2)).Return(0, fmt.Errorf("invalid data type"))
	_, err = suite.server.Incrby(suite.context, &redish.KeyIntValue{Key: stringValueKey, Value: int64(2)})
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestDecrby() {
	suite.engine.On("Decrby", intValueKey, int64(2)).Return(2, nil)
	ret, err := suite.server.Decrby(suite.context, &redish.KeyIntValue{Key: intValueKey, Value: int64(2)})
	suite.NoError(err)
	suite.Equal(int64(2), ret.Value)

	suite.engine.On("Decrby", stringValueKey, int64(2)).Return(0, fmt.Errorf("invalid data type"))
	_, err = suite.server.Decrby(suite.context, &redish.KeyIntValue{Key: stringValueKey, Value: int64(2)})
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestStrlen() {
	suite.engine.On("Strlen", stringValueKey).Return(42, nil)
	ret, err := suite.server.Strlen(suite.context, &stringValueKeyObject)
	suite.Equal(int64(42), ret.Value)
	suite.NoError(err)

	suite.engine.On("Strlen", fakeValueKey).Return(0, fmt.Errorf("invalid data type"))
	_, err = suite.server.Strlen(suite.context, &fakeValueKeyObject)
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestGetSet() {
	suite.engine.On("GetSet", stringValueKey, intValue).Return(&stringValue, nil)
	ret, err := suite.server.Getset(suite.context, &redish.KeyValue{Key: stringValueKey, Value: intValue})
	suite.Equal(stringValue, ret.Value.Value)
	suite.NoError(err)

	suite.engine.On("GetSet", doesNotExistKey, intValue).Return(nil, nil)
	ret, err = suite.server.Getset(suite.context, &redish.KeyValue{Key: doesNotExistKey, Value: intValue})
	suite.Nil(ret.Value)
	suite.NoError(err)

	suite.engine.On("GetSet", fakeValueKey, stringValue).Return(nil, fmt.Errorf("data type error"))
	_, err = suite.server.Getset(suite.context, &redish.KeyValue{Key: fakeValueKey, Value: stringValue})
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestMGet() {
	suite.engine.On("MGet", []string{stringValueKey, doesNotExistKey, fakeValueKey, intValueKey}).Return([]*string{&stringValue, nil, nil, &intValue}, nil)
	ret, err := suite.server.Mget(suite.context, &redish.KeyList{Keys: []*redish.Key{&stringValueKeyObject, &doesNotExistKeyObject, &fakeValueKeyObject, &intValueKeyObject}})
	suite.NoError(err)
	suite.Equal(stringValue, ret.Values[0].Value.Value)
	suite.Nil(ret.Values[1].Value)
	suite.Nil(ret.Values[2].Value)
	suite.Equal(intValue, ret.Values[3].Value.Value)
	suite.NoError(err)

	suite.engine.On("MGet", []string{}).Return([]*string{}, nil)
	ret, err = suite.server.Mget(suite.context, &redish.KeyList{})
	suite.NoError(err)
	suite.Equal(0, len(ret.Values))

	suite.engine.On("MGet", []string{fakeValueKey}).Return([]*string{}, fmt.Errorf("this somehow failed"))
	_, err = suite.server.Mget(suite.context, &redish.KeyList{Keys: []*redish.Key{&fakeValueKeyObject}})
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestMGetNilReturn() {
	// I don't think this is actually possible, but I want to cover that edge case!
	// and of course, this broke! this is 100% why you write tests!
	// and because this is a weird case, what do? Originally I wanted it to not error but just return an empty list, but
	// I think in this case, since we're talking about integration with an engine, we need to guard against the engine
	// doing something bad, like returning a number of values that isn't the same as the number of keys, so there should be an error
	// returned to the client. Boom.
	suite.engine.On("MGet", []string{intValueKey}).Return(nil, nil)
	_, err := suite.server.Mget(suite.context, &redish.KeyList{Keys: []*redish.Key{&intValueKeyObject}})
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestMSet() {
	// happy case
	suite.engine.On("MSet", map[string]string{stringValueKey: stringValue, intValueKey: intValue}).Return(nil)
	ok, err := suite.server.Mset(suite.context, &redish.KeyValueList{Pairs: []*redish.KeyValue{&stringKeyValue, &intKeyValue}})
	suite.Equal(&redish.OK{}, ok)
	suite.NoError(err)

	// duplicate key
	stringIntKeyValue := redish.KeyValue{Key: stringValueKey, Value: intValue}
	suite.engine.On("MSet", map[string]string{stringValueKey: intValue, intValueKey: intValue}).Return(nil)
	ok, err = suite.server.Mset(suite.context, &redish.KeyValueList{Pairs: []*redish.KeyValue{&stringKeyValue, &intKeyValue, &stringIntKeyValue}})
	suite.Equal(&redish.OK{}, ok)
	suite.NoError(err)

	// engine failure
	suite.engine.On("MSet", map[string]string{stringValueKey: intValue}).Return(fmt.Errorf("this somehow failed"))
	_, err = suite.server.Mset(suite.context, &redish.KeyValueList{Pairs: []*redish.KeyValue{&stringIntKeyValue}})
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestType() {
	suite.engine.On("Type", stringValueKey).Return("string", nil)
	ret, err := suite.server.Type(suite.context, &stringValueKeyObject)
	suite.NoError(err)
	suite.Equal("string", ret.Value.Value)

	suite.engine.On("Type", fakeValueKey).Return("fake", fmt.Errorf("this somehow failed"))
	_, err = suite.server.Type(suite.context, &fakeValueKeyObject)
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestExpire() {
	suite.engine.On("Expire", stringValueKey, int64(42)).Return(true, nil)
	ret, err := suite.server.Expire(suite.context, &redish.KeyIntValue{Key: stringValueKey, Value: int64(42)})
	suite.Equal(int64(1), ret.Value)
	suite.NoError(err)

	suite.engine.On("Expire", intValueKey, int64(42)).Return(false, nil)
	ret, err = suite.server.Expire(suite.context, &redish.KeyIntValue{Key: intValueKey, Value: int64(42)})
	suite.Equal(int64(0), ret.Value)
	suite.NoError(err)

	suite.engine.On("Expire", fakeValueKey, int64(42)).Return(false, fmt.Errorf("this somehow failed"))
	_, err = suite.server.Expire(suite.context, &redish.KeyIntValue{Key: fakeValueKey, Value: int64(42)})
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestPExpire() {
	suite.engine.On("PExpire", stringValueKey, int64(42)).Return(true, nil)
	ret, err := suite.server.Pexpire(suite.context, &redish.KeyIntValue{Key: stringValueKey, Value: int64(42)})
	suite.Equal(int64(1), ret.Value)
	suite.NoError(err)

	suite.engine.On("PExpire", intValueKey, int64(42)).Return(false, nil)
	ret, err = suite.server.Pexpire(suite.context, &redish.KeyIntValue{Key: intValueKey, Value: int64(42)})
	suite.Equal(int64(0), ret.Value)
	suite.NoError(err)

	suite.engine.On("PExpire", fakeValueKey, int64(42)).Return(false, fmt.Errorf("this somehow failed"))
	_, err = suite.server.Pexpire(suite.context, &redish.KeyIntValue{Key: fakeValueKey, Value: int64(42)})
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestExpireAt() {
	suite.engine.On("ExpireAt", stringValueKey, int64(42)).Return(true, nil)
	ret, err := suite.server.Expireat(suite.context, &redish.KeyIntValue{Key: stringValueKey, Value: int64(42)})
	suite.Equal(int64(1), ret.Value)
	suite.NoError(err)

	suite.engine.On("ExpireAt", intValueKey, int64(42)).Return(false, nil)
	ret, err = suite.server.Expireat(suite.context, &redish.KeyIntValue{Key: intValueKey, Value: int64(42)})
	suite.Equal(int64(0), ret.Value)
	suite.NoError(err)

	suite.engine.On("ExpireAt", fakeValueKey, int64(42)).Return(false, fmt.Errorf("this somehow failed"))
	_, err = suite.server.Expireat(suite.context, &redish.KeyIntValue{Key: fakeValueKey, Value: int64(42)})
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestPExpireAt() {
	suite.engine.On("PExpireAt", stringValueKey, int64(42)).Return(true, nil)
	ret, err := suite.server.Pexpireat(suite.context, &redish.KeyIntValue{Key: stringValueKey, Value: int64(42)})
	suite.Equal(int64(1), ret.Value)
	suite.NoError(err)

	suite.engine.On("PExpireAt", intValueKey, int64(42)).Return(false, nil)
	ret, err = suite.server.Pexpireat(suite.context, &redish.KeyIntValue{Key: intValueKey, Value: int64(42)})
	suite.Equal(int64(0), ret.Value)
	suite.NoError(err)

	suite.engine.On("PExpireAt", fakeValueKey, int64(42)).Return(false, fmt.Errorf("this somehow failed"))
	_, err = suite.server.Pexpireat(suite.context, &redish.KeyIntValue{Key: fakeValueKey, Value: int64(42)})
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestPersist() {
	suite.engine.On("Persist", stringValueKey).Return(true, nil)
	ret, err := suite.server.Persist(suite.context, &stringValueKeyObject)
	suite.Equal(int64(1), ret.Value)
	suite.NoError(err)

	suite.engine.On("Persist", intValueKey).Return(false, nil)
	ret, err = suite.server.Persist(suite.context, &intValueKeyObject)
	suite.Equal(int64(0), ret.Value)
	suite.NoError(err)

	suite.engine.On("Persist", doesNotExistKey).Return(false, nil)
	ret, err = suite.server.Persist(suite.context, &doesNotExistKeyObject)
	suite.Equal(int64(0), ret.Value)
	suite.NoError(err)

	suite.engine.On("Persist", fakeValueKey).Return(false, fmt.Errorf("this somehow failed"))
	_, err = suite.server.Persist(suite.context, &fakeValueKeyObject)
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestTTL() {
	suite.engine.On("TTL", stringValueKey).Return(42, nil)
	ret, err := suite.server.Ttl(suite.context, &stringValueKeyObject)
	suite.Equal(int64(42), ret.Value)
	suite.NoError(err)

	suite.engine.On("TTL", intValueKey).Return(-1, nil)
	ret, err = suite.server.Ttl(suite.context, &intValueKeyObject)
	suite.Equal(int64(-1), ret.Value)
	suite.NoError(err)

	suite.engine.On("TTL", doesNotExistKey).Return(-2, nil)
	ret, err = suite.server.Ttl(suite.context, &doesNotExistKeyObject)
	suite.Equal(int64(-2), ret.Value)
	suite.NoError(err)

	suite.engine.On("TTL", fakeValueKey).Return(0, fmt.Errorf("this somehow failed"))
	_, err = suite.server.Ttl(suite.context, &fakeValueKeyObject)
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func (suite *serverTestSuite) TestPTTL() {
	suite.engine.On("PTTL", stringValueKey).Return(42, nil)
	ret, err := suite.server.Pttl(suite.context, &stringValueKeyObject)
	suite.Equal(int64(42), ret.Value)
	suite.NoError(err)

	suite.engine.On("PTTL", intValueKey).Return(-1, nil)
	ret, err = suite.server.Pttl(suite.context, &intValueKeyObject)
	suite.Equal(int64(-1), ret.Value)
	suite.NoError(err)

	suite.engine.On("PTTL", doesNotExistKey).Return(-2, nil)
	ret, err = suite.server.Pttl(suite.context, &doesNotExistKeyObject)
	suite.Equal(int64(-2), ret.Value)
	suite.NoError(err)

	suite.engine.On("PTTL", fakeValueKey).Return(0, fmt.Errorf("this somehow failed"))
	_, err = suite.server.Pttl(suite.context, &fakeValueKeyObject)
	suite.Error(err)

	suite.engine.AssertExpectations(suite.T())
}

func TestServerTests(t *testing.T) {
	suite.Run(t, new(serverTestSuite))
}
