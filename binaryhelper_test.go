package binaryhelper

import (
	"reflect"
	"testing"
)

func TestCopyFixed(t *testing.T) {
	var i1 int32
	type args struct {
		src any
		dst any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    any
	}{
		{name: "00",
			args:    args{},
			wantErr: true,
		},
		{name: "01",
			args:    args{dst: testing.T{}},
			wantErr: true,
		},
		{name: "02",
			args:    args{src: testing.B{}},
			wantErr: true,
		},
		{name: "03",
			args:    args{src: testing.F{}, dst: testing.M{}},
			wantErr: true,
		},
		{name: "04",
			args:    args{src: [1]byte{1}, dst: [1]byte{}},
			wantErr: true,
		},
		{name: "05",
			args:    args{src: [1]byte{1}, dst: &([2]byte{})},
			wantErr: true,
		},
		{name: "1",
			args: args{src: int32(1234), dst: &i1},
			want: int32(1234),
		},
		{name: "2",
			args: args{src: [2]byte{1, 2}, dst: &([2]byte{})},
			want: [2]byte{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyFixed(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("CopyFixed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			d := reflect.Indirect(reflect.ValueOf(tt.args.dst)).Interface()
			if !reflect.DeepEqual(d, tt.want) {
				t.Errorf("CopyFixed() = %v, want %v", d, tt.want)
			}
		})
	}
}
