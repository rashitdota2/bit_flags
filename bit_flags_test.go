package bit_flags

import (
	"reflect"
	"testing"
)

const (
	_ uint8 = iota
	CreatePermission
	RemovePermission
	UpdatePermission
)

func TestHasFlag(t *testing.T) {
	type args struct {
		bitFlags BitFlags
		flag     uint8
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"HASN'T FLAG",
			args{
				bitFlags: NewBitFlags(),
				flag:     RemovePermission,
			},
			false,
		},
		{
			"HAS FLAG",
			args{
				bitFlags: map[uint8]uint64{1: 2},
				flag:     CreatePermission,
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasFlag(tt.args.bitFlags, tt.args.flag); got != tt.want {
				t.Errorf("HasFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveFlags(t *testing.T) {
	type args struct {
		bitFlags BitFlags
		flags    []uint8
	}
	tests := []struct {
		name string
		args args
		want map[uint8]uint64
	}{
		{
			"REMOVE FLAGS",
			args{
				bitFlags: map[uint8]uint64{1: 14, 2: 2},
				flags:    []uint8{CreatePermission, RemovePermission, 64},
			},
			map[uint8]uint64{1: 8, 2: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFlags(tt.args.bitFlags, tt.args.flags...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetFlags(t *testing.T) {
	type args struct {
		flagBits BitFlags
		flags    []uint8
	}
	tests := []struct {
		name string
		args args
		want map[uint8]uint64
	}{
		{
			"SET FLAGS",
			args{
				flagBits: NewBitFlags(),
				flags:    []uint8{CreatePermission, RemovePermission, UpdatePermission, 64},
			},
			map[uint8]uint64{1: 14, 2: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetFlags(tt.args.flagBits, tt.args.flags...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
