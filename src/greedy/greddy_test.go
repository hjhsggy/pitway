package greedy

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestChargeConin(t *testing.T) {

	var expect, actual, m []int

	m = []int{25, 10, 5, 1}
	expect = []int{3, 2, 0, 4}

	actual = ChargeConin(m, 99)

	assert.Equal(t, expect, actual)

}

func TestUn(t *testing.T) {
	Unmarshal()
}
