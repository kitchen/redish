package engine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type intValueStoreTestSuite struct {
	suite.Suite
	intValue    int64
	stringValue string
	store       intValueStore
}

func (suite *intValueStoreTestSuite) SetupTest() {
	suite.intValue = 123
	suite.stringValue = fmt.Sprintf("%d", suite.intValue)
	suite.store = intValueStore{intValue: suite.intValue}
}

func (suite *intValueStoreTestSuite) TestGet() {
	suite.Equal(suite.intValue, suite.store.intValue)

	value, err := suite.store.get()
	suite.Equal(suite.stringValue, value)
	suite.NoError(err)
}

func (suite *intValueStoreTestSuite) TestGetType() {
	suite.Equal("string", suite.store.getType())
}

func (suite *intValueStoreTestSuite) TestIncrBy() {
	newExpectedValue := suite.intValue + 10
	newValue, err := suite.store.incrby(10)
	suite.Equal(newExpectedValue, newValue)
	suite.NoError(err)

	newExpectedValue -= 20
	newValue, err = suite.store.incrby(-20)
	suite.Equal(newExpectedValue, newValue)
	suite.NoError(err)
}

func TestIntValueStore(t *testing.T) {
	suite.Run(t, new(intValueStoreTestSuite))
}
