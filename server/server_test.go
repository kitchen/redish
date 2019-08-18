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
var intValue = "424242"
var intValueKey = "intvalue"
var intValueKeyObject = redish.Key{Key: intValueKey}
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

func TestServerTests(t *testing.T) {
	suite.Run(t, new(serverTestSuite))
}
