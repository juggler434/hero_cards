package helpers

import "testing"

const (
	large = 4
	small = 3
)

func TestMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "large number first",
			args: args{
				a: large,
				b: small,
			},
			want: small,
		},
		{
			name: "small number first",
			args: args{
				a: small,
				b: large,
			},
			want: small,
		},
		{
			name: "equal numbers",
			args: args{
				a: large,
				b: large,
			},
			want: large,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "small number first",
			args: args{
				a: small,
				b: large,
			},
			want: large,
		},
		{
			name: "large number first",
			args: args{
				a: large,
				b: small,
			},
			want: large,
		},
		{
			name: "equal values",
			args: args{
				a: small,
				b: small,
			},
			want: small,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
