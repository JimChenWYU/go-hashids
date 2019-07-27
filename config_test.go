package hashids

import (
	"reflect"
	"testing"
)

func TestNewHashidConfig(t *testing.T) {
	type args struct {
		_salt      string
		_minLength int
		_alphabet  string
	}
	tests := []struct {
		name string
		args args
		want *HashidsConfig
	}{
		{
			name: "Cast 1",
			args: args{
				_salt:      "Cast",
				_minLength: 2,
				_alphabet:  "abcdefghijklmnopqrst",
			},
			want: &HashidsConfig{
				minLength:        2,
				originalSalt:     "Cast",
				originalAlphabet: "abcdefghijklmnopqrst",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHashidConfig(tt.args._salt, tt.args._minLength, tt.args._alphabet)
			if reflect.ValueOf(got).Type() != reflect.ValueOf(tt.want).Type() {
				t.Error("got object not equal type to want")
			}
			if got.originalAlphabet != tt.want.originalAlphabet {
				t.Errorf("originalAlphabet = %v, want %v", got, tt.want.originalAlphabet)
			}
			if got.originalAlphabet != tt.want.originalAlphabet {
				t.Errorf("originalAlphabet = %v, want %v", got, tt.want.originalAlphabet)
			}
		})
	}
	t.Run("Cast 2", func(t *testing.T) {
		cast := args{
			_salt:      "",
			_minLength: 0,
			_alphabet:  "abc",
		}
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("should receive a panic.")
			}
		}()

		NewHashidConfig(cast._salt, cast._minLength, cast._alphabet)
	})

	t.Run("Cast 3", func(t *testing.T) {
		cast := args{
			_salt:      "bar",
			_minLength: 10,
			_alphabet:  "abcdefghijklmnopq",
		}

		config := NewHashidConfig(cast._salt, cast._minLength, cast._alphabet)
		tmp := *config
		_config := &tmp

		config.notify("foo", 8, "qwertyuiopasdfghj")

		if string(config.salt) == string(_config.salt) {
			t.Errorf("salt = %v should be not equal to original salt = %v.", string(config.salt), string(_config.salt))
		}
		if string(config.alphabet) == string(_config.alphabet) {
			t.Errorf("alphabet = %v should be not equal to original alphabet = %v.", string(config.alphabet), string(_config.alphabet))
		}
		t.Logf("config = %v | _config = %v", config, _config)
	})
}
