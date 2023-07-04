package colors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	count, err := Count("../../data", "blue")
	assert.NoError(t, err)
	assert.Equal(t, 1000, count)
}
