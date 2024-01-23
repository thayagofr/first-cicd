package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("", func(t *testing.T) {
		numbersToSum := []int{0, 0, 1}

		got := Sum(numbersToSum)
		assert.Equal(t, 1, got)
	})

	t.Run("", func(t *testing.T) {
		numbersToSum := []float64{0, 2.2, 1.1}

		got := Sum(numbersToSum)
		assert.InDelta(t, 3.3, got, 0.01)
	})
}
