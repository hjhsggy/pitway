package dynamic

import (
	"testing"
)

func TestMatchRegex(t *testing.T) {

}

func TestFib(t *testing.T) {

	var dest int = 20

	num := Fib(dest)

	t.Logf("Fib(%v)=%v\n", dest, num)
}

func TestMemoFib(t *testing.T) {

	var dest int = 20

	num := FibWithMemo(dest)

	t.Logf("Fib(%v)=%v\n", dest, num)

}

func TestDBFit(t *testing.T) {

	var dest int = 20

	num := FibWithDP(dest)

	t.Logf("Fib(%v)=%v\n", dest, num)

}

func TestCoinChange(t *testing.T) {

	var dest int = 5

	num := CoinChange(dest)

	t.Logf("CoinChange(%v)=%v\n", dest, num)

}

func TestCoinMun(t *testing.T) {

	CoinNums(24)

}

func TestSubString(t *testing.T) {

	txt := "ccvgfvgt"
	pat := "vgt"

	result := CommSubStr(txt, pat)
	t.Log(result)

	//assert.Equal(t, 5, result)

}

func TestGetNext(t *testing.T) {

	patStr := []string{
		"abcvgfvg",
		"ababcabb",
		"abcbcbab",
		"ababaca",
	}

	for _, val := range patStr {

		t.Log(substringRepetition(val))

	}

}

func TestKMP(t *testing.T) {

	s := "ABD"
	p := "ABABCABD"

	db := KMP(s)

	t.Log(db)

	t.Log(searchKMP(s, p))

}
