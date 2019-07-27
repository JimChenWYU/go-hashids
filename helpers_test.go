package hashids

import (
	"reflect"
	"testing"
)

func Test_uniqueCharacter(t *testing.T) {
	type args struct {
		alphabet string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TestCase
		{
			name: "Nerver Repeat",
			args: args{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"},
			want: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		},
		{
			name: "Case 'aaaaa'",
			args: args{alphabet: "aaaaa"},
			want: "a",
		},
		{
			name: "Case 'abcabc'",
			args: args{alphabet: "abcabc"},
			want: "abc",
		},
		{
			name: "Case 'aabcc'",
			args: args{alphabet: "aabcc"},
			want: "abc",
		},
		{
			name: "Case 'a_b_c'",
			args: args{alphabet: "a_b_c"},
			want: "a_bc",
		},
		{
			name: "Case 'a101023020'",
			args: args{alphabet: "a101023020"},
			want: "a1023",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueCharacter(tt.args.alphabet); got != tt.want {
				t.Errorf("uniqueCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeSpaces(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Remove First Space",
			args: args{str: " abc"},
			want: "abc",
		},
		{
			name: "Remove Center Space",
			args: args{str: "a b c"},
			want: "abc",
		},
		{
			name: "Remove Last Space",
			args: args{str: "abc "},
			want: "abc",
		},
		{
			name: "Remove Multi Spaces",
			args: args{str: "a   b  c"},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSpaces(tt.args.str); got != tt.want {
				t.Errorf("removeSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_consistentShuffle(t *testing.T) {
	alphabet := []rune("foobar")
	type args struct {
		alphabet []rune
		salt     []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "Empty salt.",
			args: args{
				alphabet: alphabet,
				salt:     []rune{},
			},
			want: alphabet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := consistentShuffle(tt.args.alphabet, tt.args.salt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("consistentShuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
