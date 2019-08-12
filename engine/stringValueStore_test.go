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

func (suite *stringValueStoreTestSuite) SetupTests() {
	suite.stringValue = "abc"
	suite.store = stringValueStore{stringValue: suite.stringValue}
}

func (suite *stringValueStoreTestSuite) TestGet() {
	suite.Equal(suite.stringValue, suite.store.stringValue)
	stringValue, err := suite.store.get()
	suite.Equal(suite.stringValue, stringValue)
	suite.NoError(err)
}

func (suite *stringValueStoreTestSuite) TestGetType() {
	suite.Equal("string", suite.store.getType())
}

func (suite *stringValueStoreTestSuite) TestIncrBy() {
	_, err := suite.store.incrby(1)
	suite.Error(err)
}

func TestStringValueStore(t *testing.T) {
	suite.Run(t, new(stringValueStoreTestSuite))
}
