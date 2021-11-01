package main

import (
	"testing"
)

func Test_findByPk(t *testing.T) {

	res := myAtoi("3.1456")
	t.Log(res == 3)
	t.Log(res)

	res = myAtoi("420")
	t.Log(res == 420)
	t.Log(res)

	res = myAtoi("   -42")
	t.Log(res == -42)
	t.Log(res)

	res = myAtoi("4193 with words")
	t.Log(res == 4193)
	t.Log(res)

	res = myAtoi("words and 987")
	t.Log(res == 0)
	t.Log(res)

	res = myAtoi("-91283472332")
	t.Log(res == -2147483648)
	t.Log(res)
}
