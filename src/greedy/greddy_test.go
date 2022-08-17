package greedy

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestChargeCoin(t *testing.T) {

	var expect, actual, m []int

	m = []int{25, 10, 5, 1}
	expect = []int{3, 2, 0, 4}

	actual = ChargeCoin(m, 99)

	assert.Equal(t, expect, actual)

}

func TestUn(t *testing.T) {
	Unmarshal()
}

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
		})
	}
}
