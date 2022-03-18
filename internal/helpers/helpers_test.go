package helpers

import "testing"

func TestParseToFloat(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "SuccessFloat",
			args: args{
				str: "123.9",
			},
			want: 123.9,
		},
		{
			name: "SuccessCeil",
			args: args{
				str: "123",
			},
			want: 123,
		},
		{
			name: "FailFloat",
			args: args{
				str: "123abc",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseToFloat(tt.args.str); got != tt.want {
				t.Errorf("ParseToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
