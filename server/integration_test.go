package main

import (
	"context"
	"testing"

	"github.com/kitchen/redish/engine"
	"github.com/kitchen/redish/grpc_server"
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

type redishServerIntegrationTestSuite struct {
	engine  engine.Engine
	server  redish.RedishServer
	context context.Context
	suite.Suite
}

func (suite *redishServerIntegrationTestSuite) SetupTest() {
	suite.engine = engine.NewEngine()
	suite.server = grpc_server.NewServer(suite.engine)
	suite.context = context.TODO()
}

func (suite *redishServerIntegrationTestSuite) TestSetGet() {
	ok, err := suite.server.Set(suite.context, &redish.SetRequest{Key: stringValueKey, Value: stringValue})
	suite.Equal(&redish.OK{}, ok)
	suite.NoError(err)

	ret, err := suite.server.Get(suite.context, &stringValueKeyObject)
	suite.NoError(err)
	suite.Equal(stringValue, ret.Value.Value)

	ok, err = suite.server.Set(suite.context, &redish.SetRequest{Key: intValueKey, Value: intValue})
	suite.Equal(&redish.OK{}, ok)
	suite.NoError(err)

	ret, err = suite.server.Get(suite.context, &intValueKeyObject)
	suite.Equal(intValue, ret.Value.Value)
	suite.NoError(err)

	ret, err = suite.server.Get(suite.context, &doesNotExistKeyObject)
	suite.Nil(ret.Value)
	suite.NoError(err)
}

func TestRunIntegrationTests(t *testing.T) {
	suite.Run(t, new(redishServerIntegrationTestSuite))
}
