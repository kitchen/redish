package engine

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type setValueStoreTestSuite struct {
	store setValueStore
	suite.Suite
}

func (suite *setValueStoreTestSuite) SetupTest() {
	suite.store = setValueStore{}
}

func (suite *setValueStoreTestSuite) TestAdd() {

}

func TestRunSetValueStoreTests(t *testing.T) {
	suite.Run(t, new(setValueStoreTestSuite))
}
