package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	hasil := Solution([]string{"test1a", "test2", "test1b", "test1c", "test3"},
		[]string{"wrong", "OK", "runtime", "OK", "time exceed"})

	assert.Equal(t, 33, hasil)
}
