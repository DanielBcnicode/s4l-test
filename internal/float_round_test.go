package internal

import "testing"

func TestFloatRoundPrecision(t *testing.T) {
	type args struct {
		n float32
		p int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "One decimal",
			args: args{
				n: 3.1416,
				p: 1,
			},
			want: 3.1,
		},
		{
			name: "Two decimals",
			args: args{
				n: 3.1416,
				p: 2,
			},
			want: 3.14,
		},
		{
			name: "Three decimals",
			args: args{
				n: 3.1416,
				p: 3,
			},
			want: 3.142,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatRoundPrecision(tt.args.n, tt.args.p); got != tt.want {
				t.Errorf("FloatRoundPrecision() = %v, want %v", got, tt.want)
			}
		})
	}
}
