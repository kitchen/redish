package engine

import (
	"testing"

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

func (suite *valueStoreTestSuite) TestDefaultFailures() {
	_, err := suite.store.get()
	suite.Error(err)

	_, err = suite.store.incrby(1)
	suite.Error(err)
}

func (suite *valueStoreTestSuite) TestInterface() {
	var store valueStoreInterface
	store = &fakeValueStore{}
	suite.NotNil(store)
}

func (suite *valueStoreTestSuite) TestDefaultPasses() {
	suite.Equal("fake", suite.store.getType())
}

func TestValueStoreTestSuite(t *testing.T) {
	suite.Run(t, new(valueStoreTestSuite))
}

//
// func TestDefaultFailures(t *testing.T) {
// 	store := fakeValueStore{}
// 	_, err := store.get()
// 	assert.Error(t, err)
//
// 	_, err = store.incrby(1)
// 	assert.Error(t, err)
//
// 	assert.Equal(t, "fake", store.getType())
// }
