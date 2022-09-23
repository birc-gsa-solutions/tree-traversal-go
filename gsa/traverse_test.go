package gsa

import (
	"reflect"
	"testing"
)

func TestInOrder(t *testing.T) {
	type args struct {
		v *T
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"Test 1",
			args{&T{2, &T{1, nil, nil}, &T{4, &T{3, nil, nil}, &T{5, nil, nil}}}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"Test 2",
			args{nil},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InOrder(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInOrder2(t *testing.T) {
	type args struct {
		v *T
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"Test 1",
			args{&T{2, &T{1, nil, nil}, &T{4, &T{3, nil, nil}, &T{5, nil, nil}}}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"Test 2",
			args{nil},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InOrder2(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBfOrder(t *testing.T) {
	type args struct {
		v *T
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"Test 1",
			args{&T{2, &T{1, nil, nil}, &T{4, &T{3, nil, nil}, &T{5, nil, nil}}}},
			[]int{2, 1, 4, 3, 5},
		},
		{
			"Test 2",
			args{nil},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BfOrder(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BfOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
