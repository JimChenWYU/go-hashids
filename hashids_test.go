package hashids

import (
	"reflect"
	"testing"
)

var options = NewHashidConfig("this is my salt", 30, DEFAULT_ALPHABET)

func TestHashids_EncodeInt64(t *testing.T) {
	type fields struct {
		options *HashidsConfig
	}
	type args struct {
		numbers []int64
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "empty array of numbers",
			fields:  fields{options: options},
			args:    args{numbers: []int64{}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Cast 60",
			fields:  fields{options: options},
			args:    args{numbers: []int64{60}},
			wantErr: false,
			want:    "xo3XwO8JMz2bVGn2zG4D6vQRak9ZrN",
		},
		{
			name:    "Cast 1,2,3",
			fields:  fields{options: options},
			args:    args{numbers: []int64{1, 2, 3}},
			wantErr: false,
			want:    "ZPVgxzNb59LGlaHquq06DmlyMX3okO",
		},
		{
			name:    "Cast 1,2,9",
			fields:  fields{options: options},
			args:    args{numbers: []int64{1, 2, 9}},
			wantErr: false,
			want:    "xz6JjalokwMA4DHOueA2gOWXe4b87y",
		},
		{
			name:    "Cast 26, 27, 28",
			fields:  fields{options: options},
			args:    args{numbers: []int64{26, 27, 28}},
			wantErr: false,
			want:    "4KjXVoQMzD8AWRfySx06LYgl29OvP3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Hashids{
				options: tt.fields.options,
			}
			got, err := this.EncodeInt64(tt.args.numbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashids_DecodeInt64(t *testing.T) {
	type fields struct {
		options *HashidsConfig
	}
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int64
		wantErr bool
	}{
		{
			name:    "Empty string",
			fields:  fields{options: options},
			args:    args{hash: ""},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Case xo3XwO8JMz2bVGn2zG4D6vQRak9ZrN to [60]",
			fields:  fields{options: options},
			args:    args{hash: "xo3XwO8JMz2bVGn2zG4D6vQRak9ZrN"},
			want:    []int64{60},
			wantErr: false,
		},
		{
			name:    "Case ZPVgxzNb59LGlaHquq06DmlyMX3okO to [1,2,3]",
			fields:  fields{options: options},
			args:    args{hash: "ZPVgxzNb59LGlaHquq06DmlyMX3okO"},
			want:    []int64{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "Cast xz6JjalokwMA4DHOueA2gOWXe4b87y to [1,2,9]",
			fields:  fields{options: options},
			args:    args{hash: "xz6JjalokwMA4DHOueA2gOWXe4b87y"},
			want:    []int64{1, 2, 9},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Hashids{
				options: tt.fields.options,
			}
			got, err := this.DecodeInt64(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashids_EncodeHexWithError(t *testing.T) {
	type fields struct {
		options *HashidsConfig
	}
	type args struct {
		hex string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Case abc encode to 1kNW4RKP9Xzjd5RzWGvOElgbLeqVmD",
			fields:  fields{options: options},
			args:    args{hex: "abc"},
			want:    "1kNW4RKP9Xzjd5RzWGvOElgbLeqVmD",
			wantErr: false,
		},
		{
			name:    "Case 18af encode to Woz1Dw4BvXmRGkZQLde9M7k2rK63Yp",
			fields:  fields{options: options},
			args:    args{hex: "18af"},
			want:    "Woz1Dw4BvXmRGkZQLde9M7k2rK63Yp",
			wantErr: false,
		},
		{
			name:    "Case 9876 encode to v6zyb5ZYpmVLdKVZz4A7x89gjnPNq2",
			fields:  fields{options: options},
			args:    args{hex: "9876"},
			want:    "v6zyb5ZYpmVLdKVZz4A7x89gjnPNq2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Hashids{
				options: tt.fields.options,
			}
			got, err := this.EncodeHexWithError(tt.args.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeHexWithError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeHexWithError() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashids_DecodeHexWithError(t *testing.T) {
	type fields struct {
		options *HashidsConfig
	}
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Case 2nmvd5n7EOQ16bDFPBvplDgLqGkqYR decode to abcabcabcabcabcabcabcabc",
			fields:  fields{options: options},
			args:    args{hash: "2nmvd5n7EOQ16bDFPBvplDgLqGkqYR"},
			want:    "abcabcabcabcabcabcabcabc",
			wantErr: false,
		},
		{
			name:    "Case vLw35q21ylPAZLob68GMJepzEQZ4XR decode to 123456",
			fields:  fields{options: options},
			args:    args{hash: "vLw35q21ylPAZLob68GMJepzEQZ4XR"},
			want:    "123456",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Hashids{
				options: tt.fields.options,
			}
			got, err := this.DecodeHexWithError(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeHexWithError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeHexWithError() got = %v, want %v", got, tt.want)
			}
		})
	}
}
