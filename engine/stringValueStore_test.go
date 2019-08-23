package engine

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type stringValueStoreTestSuite struct {
	suite.Suite
	stringValue string
	store       stringValueStore
}

func (suite *stringValueStoreTestSuite) SetupTest() {
	suite.stringValue = "abc"
	suite.store = stringValueStore{stringValue: suite.stringValue}
}

func (suite *stringValueStoreTestSuite) TestGet() {
	suite.Equal("abc", suite.stringValue)
	suite.Equal(suite.stringValue, suite.store.stringValue)
	stringValue := suite.store.get()
	suite.Equal(suite.stringValue, stringValue)
}

func (suite *stringValueStoreTestSuite) TestGetType() {
	suite.Equal("string", suite.store.getType())
}

func (suite *stringValueStoreTestSuite) TestIncrBy() {
	_, err := suite.store.incrby(1)
	suite.Error(err)
}

func (suite *stringValueStoreTestSuite) TestImplements() {
	var baseInterface *valueStoreInterface
	suite.Implements(baseInterface, &suite.store)

	var stringishInterface *stringishValueStoreInterface
	suite.Implements(stringishInterface, &suite.store)
}

func TestStringValueStore(t *testing.T) {
	suite.Run(t, new(stringValueStoreTestSuite))
}
