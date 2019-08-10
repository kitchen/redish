package engine

import "testing"
import "github.com/stretchr/testify/assert"

func TestEngineIncr(t *testing.T) {
	engine := NewEngine()
	value, err := engine.Incr("abc")
	assert.Equal(t, int64(1), *value)
	assert.NoError(t, err)
}
