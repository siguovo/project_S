package stools

import (
	"testing"
)

func TestShowVersion(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			want: "stools_v1.1.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShowVersion(); got != tt.want {
				t.Errorf("ShowVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSAdd(t *testing.T) {
	type args struct {
		numA int
		numB int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "testSAdd1",
			args: args{
				10,
				10,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SAdd(tt.args.numA, tt.args.numB); got != tt.want {
				t.Errorf("SAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
