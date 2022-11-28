package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplyStrin(t *testing.T) {
	t.Log("Test: multiply string")
	t.Run("Example 1", func(t *testing.T) {
		num1 := "2"
		num2 := "3"
		result := multiply(num1, num2)
		assert.Equal(t, "6", result)
		result = multiply2(num1, num2)
		assert.Equal(t, "6", result)
	})

	t.Run("Example 2", func(t *testing.T) {
		num1 := "123"
		num2 := "456"
		result := multiply(num1, num2)
		assert.Equal(t, "56088", result)
		result = multiply2(num1, num2)
		assert.Equal(t, "56088", result)
	})

	t.Run("Example 2", func(t *testing.T) {
		num1 := "456"
		num2 := "123"
		result := multiply(num1, num2)
		assert.Equal(t, "56088", result)
		result = multiply2(num1, num2)
		assert.Equal(t, "56088", result)
	})

	t.Run("Example 2", func(t *testing.T) {
		num1 := "81"
		num2 := "9"
		result := multiply(num1, num2)
		assert.Equal(t, "729", result)
		result = multiply2(num1, num2)
		assert.Equal(t, "729", result)
	})

	t.Run("Example 3", func(t *testing.T) {
		num1 := "9"
		num2 := "9"
		result := multiply(num1, num2)
		assert.Equal(t, "81", result)
		result = multiply2(num1, num2)
		assert.Equal(t, "81", result)
	})

	t.Run("Example 3", func(t *testing.T) {
		num1 := "498828660196"
		num2 := "840477629533"
		result := multiply(num1, num2)
		assert.Equal(t, "419254329864656431168468", result)
		result = multiply2(num1, num2)
		assert.Equal(t, "419254329864656431168468", result)
	})

	t.Run("Example 3", func(t *testing.T) {
		num1 := "987654321"
		num2 := "123456789"
		result := multiply(num1, num2)
		assert.Equal(t, "121932631112635269", result)
		result = multiply2(num1, num2)
		assert.Equal(t, "121932631112635269", result)
	})

	t.Run("Example 3", func(t *testing.T) {
		num1 := "123456789"
		num2 := "987654321"
		result := multiply(num1, num2)
		assert.Equal(t, "121932631112635269", result)
		result = multiply2(num1, num2)
		assert.Equal(t, "121932631112635269", result)
	})
}
