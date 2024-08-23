package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMessage(t *testing.T) {
	repository := NewHello()
	m, err := repository.GetMessage()
	require.NoError(t, err)
	assert.Equal(t, "Hello, World!", m.Message)
}
