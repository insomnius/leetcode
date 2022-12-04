package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombina(t *testing.T) {
	t.Log("Test: letter combinations")
	t.Run("Example 1", func(t *testing.T) {
		digits := "23"
		result := letterCombinations(digits)
		assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, result)
	})

	t.Run("Example 2", func(t *testing.T) {
		digits := "29"
		result := letterCombinations(digits)
		assert.Equal(t, []string{"aw", "ax", "ay", "az", "bw", "bx", "by", "bz", "cw", "cx", "cy", "cz"}, result)
	})

	t.Run("Example 3", func(t *testing.T) {
		digits := "279"
		result := letterCombinations(digits)
		assert.Equal(t, []string{"apw", "apx", "apy", "apz", "aqw", "aqx", "aqy", "aqz", "arw", "arx", "ary", "arz", "asw", "asx", "asy", "asz", "bpw", "bpx", "bpy", "bpz", "bqw", "bqx", "bqy", "bqz", "brw", "brx", "bry", "brz", "bsw", "bsx", "bsy", "bsz", "cpw", "cpx", "cpy", "cpz", "cqw", "cqx", "cqy", "cqz", "crw", "crx", "cry", "crz", "csw", "csx", "csy", "csz"}, result)
	})

	t.Run("Example 4", func(t *testing.T) {
		digits := "2"
		result := letterCombinations(digits)
		assert.Equal(t, []string{"a", "b", "c"}, result)
	})

	t.Run("Example 5", func(t *testing.T) {
		digits := "9"
		result := letterCombinations(digits)
		assert.Equal(t, []string{"w", "x", "y", "z"}, result)
	})
}
