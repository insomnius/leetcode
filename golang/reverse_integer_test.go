package leetcode

import "testing"

func TestReverseInteger(t *testing.T) {
	inputs := []int{
		321,
		120,
		-321,
		123,
		9999999999999999,
		0,
	}

	expected := []int{
		123,
		21,
		-123,
		321,
		0,
		0,
	}

	for k, i := range inputs {
		res := reverseInteger(i)

		if res != expected[k] {
			t.Errorf("%d is not %d", res, expected[k])
		}
	}

}
