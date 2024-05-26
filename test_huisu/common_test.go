package huisu_test

import (
	"testing"
)

func TestCommon(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	i := 1
	for _, v := range ints {
		if v == i {
			t.Logf("v: %d", v)
		} else {
			t.Errorf("v: %d, want: %d", v, i)
		}
		i++
	}
}
