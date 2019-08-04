package hashids

import (
	"errors"
	"reflect"
	"testing"
)

func TestHashids_EncodeInt64_Return_With_ERROR(t *testing.T) {
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
		wantErr error
	}{
		{
			name:    "empty array of numbers",
			fields:  fields{options: options},
			args:    args{numbers: []int64{}},
			want:    "",
			wantErr: errors.New("can not encoding empty array of numbers"),
		},
		{
			name:    "negative number",
			fields:  fields{options: options},
			args:    args{numbers: []int64{-1, -2}},
			want:    "",
			wantErr: errors.New("negative number not supported"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Hashids{
				options: tt.fields.options,
			}
			got, err := this.EncodeInt64(tt.args.numbers)
			if err == nil || err.Error() != tt.wantErr.Error() {
				t.Errorf("EncodeInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashids_DecodeInt64_Return_With_ERROR(t *testing.T) {
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
		wantErr error
	}{
		{
			name:    "unexpected hash string",
			fields:  fields{options: options},
			args:    args{hash: ""},
			want:    nil,
			wantErr: errors.New("unexpected hash string"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Hashids{
				options: tt.fields.options,
			}
			got, err := this.DecodeInt64(tt.args.hash)
			if err == nil || err.Error() != tt.wantErr.Error() {
				t.Errorf("DecodeInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashids_EncodeHex_Return_With_Error(t *testing.T) {
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
		wantErr error
	}{
		{
			name:    "Case invalid hex digit",
			fields:  fields{options: options},
			args:    args{hex: "poiuy"},
			want:    "",
			wantErr: errors.New("invalid hex digit"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Hashids{
				options: tt.fields.options,
			}
			got, err := this.EncodeHexWithError(tt.args.hex)
			if err == nil || err.Error() != tt.wantErr.Error() {
				t.Errorf("EncodeHexWithError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeHexWithError() got = %v, want %v", got, tt.want)
			}
		})
	}
}
