package padding

import (
	"math"
	"testing"
)

func TestPad(t *testing.T) {
	const PadInvalid PadType = math.MaxInt64

	type args struct {
		source  string
		length  int
		padStr  string
		padType PadType
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "pad on the left",
			args: args{
				source:  "hello",
				length:  10,
				padStr:  "!",
				padType: PadLeft,
			},
			want:    "!!!!!hello",
			wantErr: false,
		},
		{
			name: "pad on the right",
			args: args{
				source:  "hello",
				length:  10,
				padStr:  "!",
				padType: PadRight,
			},
			want:    "hello!!!!!",
			wantErr: false,
		},
		{
			name: "source is equal to the specified length",
			args: args{
				source:  "hello",
				length:  5,
				padStr:  "!",
				padType: PadRight,
			},
			want:    "hello",
			wantErr: false,
		},
		{
			name: "need trim left side",
			args: args{
				source:  "hello",
				length:  10,
				padStr:  "!!!",
				padType: PadLeft,
			},
			want:    "!!!!!hello",
			wantErr: false,
		},
		{
			name: "need trim right side",
			args: args{
				source:  "hello",
				length:  10,
				padStr:  "!!!",
				padType: PadRight,
			},
			want:    "hello!!!!!",
			wantErr: false,
		},
		{
			name: "invalid padding type",
			args: args{
				source:  "hello",
				length:  10,
				padStr:  "!",
				padType: PadInvalid,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pad(tt.args.source, tt.args.length, tt.args.padStr, tt.args.padType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAlreadyReachedLength(t *testing.T) {
	type args struct {
		source string
		length int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "source is shorter than the specified length",
			args: args{
				source: "str",
				length: 10,
			},
			want: false,
		},
		{
			name: "source is longer than the specified length",
			args: args{
				source: "str",
				length: 2,
			},
			want: true,
		},
		{
			name: "source is equal to the specified length",
			args: args{
				source: "str",
				length: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlreadyReachedLength(tt.args.source, tt.args.length); got != tt.want {
				t.Errorf("isAlreadyReachedLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
