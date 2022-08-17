package practice

import "testing"

func Test_is2Power(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{9},
			want: false,
		},
		{
			args: args{8},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := is2Power(tt.args.n); got != tt.want {
				t.Errorf("is2Power() = %v, want %v", got, tt.want)
			}
		})
	}
}
