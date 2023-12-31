package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// example_add_test.go dosya adları mutlaka _test ile bitmeli
func TestAdd(t *testing.T) {
	t.Run("TestAdd", func(t *testing.T) {
		actual := Add(10, 20)
		assert.Equal(t, 30, actual)
	})
}

func Add(x int, y int) int {
	return x + y
}
