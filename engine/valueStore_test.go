package engine

import (
	"testing"

	"time"

	"github.com/stretchr/testify/suite"
)

type fakeValueStore struct {
	valueStore
}

func (s *fakeValueStore) getType() string {
	return "fake"
}

type valueStoreTestSuite struct {
	suite.Suite
	store fakeValueStore
}

func (suite *valueStoreTestSuite) SetupTest() {
	suite.store = fakeValueStore{}
}

func (suite *valueStoreTestSuite) TestInterface() {
	var store valueStoreInterface
	store = &fakeValueStore{}
	suite.NotNil(store)
}

func (suite *valueStoreTestSuite) TestDefaultPasses() {
	suite.Equal("fake", suite.store.getType())
}

func (suite *valueStoreTestSuite) TestExpirationMethods() {
	tomorrow := time.Now().Add(time.Duration(24 * time.Hour))
	yesterday := time.Now().Add(time.Duration(-24 * time.Hour))

	suite.store.expire(&tomorrow)
	suite.NotNil(suite.store.expires())
	suite.False(suite.store.expired())

	suite.store.expire(&yesterday)
	suite.NotNil(suite.store.expires())
	suite.True(suite.store.expired())

	suite.store.expire(nil)
	suite.Nil(suite.store.expires())
	suite.False(suite.store.expired())
}

func TestValueStoreTestSuite(t *testing.T) {
	suite.Run(t, new(valueStoreTestSuite))
}
