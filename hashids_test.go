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
