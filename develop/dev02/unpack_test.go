package unpack

import (
	"fmt"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    string
		wantErr error
	}{
		{
			name:    "err case два числа подряд",
			arg:     "a2b45cd",
			want:    "a2b45cd",
			wantErr: ErrIncorrectString,
		},
		{
			name:    "err case число в начале",
			arg:     "1abcd5",
			want:    "1abcd5",
			wantErr: ErrIncorrectString,
		},
		{
			name:    "case1",
			arg:     "a3bc2d5e",
			want:    "aaabccddddde",
			wantErr: nil,
		},
		{
			name:    "edge case число в конце",
			arg:     "abcd3",
			want:    "abcddd",
			wantErr: nil,
		},
		{
			name:    "case только буквы",
			arg:     "abcd",
			want:    "abcd",
			wantErr: nil,
		},
		{
			name:    "case empty",
			arg:     "",
			want:    "",
			wantErr: ErrIncorrectString,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Unpack(test.arg)
			fmt.Println("arg:", test.arg)
			fmt.Println("got:", got)
			fmt.Println("expected:", test.want)
			fmt.Println()
			if err != test.wantErr {
				t.Errorf("Unpack() error = %v, want.err %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("Unpack() = %v, want %v", got, test.want)
			}
		})
	}
}
