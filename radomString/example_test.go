package radomString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStringRunes(t *testing.T) {
	s := RandStringRunes(10)
	assert.IsType(t, "string", s)
}

func TestRandStringBytes(t *testing.T) {
	s := RandStringBytes(10)
	assert.IsType(t, "string", s)
}
