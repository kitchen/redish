package engine

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type fakeEngine struct {
	UnimplementedEngine
}

type unimplementedEngineTestSuite struct {
	engine Engine
	suite.Suite
}

func (suite *unimplementedEngineTestSuite) SetupTest() {
	suite.engine = &fakeEngine{}
}

func (suite *unimplementedEngineTestSuite) TestMethods() {
	_, err := suite.engine.Get(stringValueKey)
	suite.Error(err)

	err = suite.engine.Set(stringValueKey, stringValue)
	suite.Error(err)

	_, err = suite.engine.Del([]string{stringValueKey})
	suite.Error(err)

	_, err = suite.engine.Exists([]string{stringValueKey})
	suite.Error(err)

	_, err = suite.engine.Incr(stringValueKey)
	suite.Error(err)
	_, err = suite.engine.Decr(stringValueKey)
	suite.Error(err)
	_, err = suite.engine.Incrby(stringValueKey, int64(0))
	suite.Error(err)
	_, err = suite.engine.Decrby(stringValueKey, int64(0))
	suite.Error(err)

	_, err = suite.engine.Strlen(stringValueKey)
	suite.Error(err)

	_, err = suite.engine.GetSet(stringValueKey, intValueString)
	suite.Error(err)

	_, err = suite.engine.MGet([]string{stringValueKey})
	suite.Error(err)

	err = suite.engine.MSet(map[string]string{stringValueKey: intValueString})
	suite.Error(err)

	_, err = suite.engine.Type(stringValueKey)
	suite.Error(err)

	_, err = suite.engine.Expire(stringValueKey, int64(100))
	suite.Error(err)
	_, err = suite.engine.PExpire(stringValueKey, int64(100))
	suite.Error(err)
	_, err = suite.engine.ExpireAt(stringValueKey, int64(100))
	suite.Error(err)
	_, err = suite.engine.PExpireAt(stringValueKey, int64(100))
	suite.Error(err)
	_, err = suite.engine.Persist(stringValueKey)
	suite.Error(err)
	_, err = suite.engine.TTL(stringValueKey)
	suite.Error(err)
	_, err = suite.engine.PTTL(stringValueKey)
	suite.Error(err)
}

func TestUnimplementedEngine(t *testing.T) {
	suite.Run(t, new(unimplementedEngineTestSuite))
}
