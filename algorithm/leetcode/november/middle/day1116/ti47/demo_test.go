package ti47

import "testing"

func TestDemo(t *testing.T) {
	unique := permuteUnique([]int{1, 1, 2})
	t.Log(unique)
}
