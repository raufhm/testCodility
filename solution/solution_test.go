package solution

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSolution(t *testing.T) {
	hasil := Solution([]string{"test1a", "test2", "test1b", "test1c", "test3"},
		[]string{"wrong", "OK", "runtime", "OK", "time exceed"})

	assert.Equal(t, 33, hasil)
}

func TestGetLen(t *testing.T) {
	result := GetLen("test1a")
	_ = result
}

func TestSolutionTwo(t *testing.T) {
	_ = SolutionTwo([]string{"test1a", "test2", "test1b", "test1c", "test3"},
		[]string{"wrong", "OK", "runtime", "OK", "time exceed"})
}
