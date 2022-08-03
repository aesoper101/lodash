package arrayx

import (
	"reflect"
	"testing"
)

func TestTail(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{arr: []int{1, 2, 3, 4, 5}},
			want: []int{2, 3, 4, 5},
		},
		{
			args: args{
				arr: []int{1},
			},
			want: []int{},
		},
		{
			args: args{
				arr: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tail(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTake(t *testing.T) {
	type args struct {
		array []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				array: []int{1, 2, 3},
				n:     1,
			},
			want: []int{1},
		},
		{
			args: args{
				array: []int{1, 2, 3},
				n:     2,
			},
			want: []int{1, 2},
		},
		{
			args: args{
				array: []int{1, 2, 3},
				n:     5,
			},
			want: []int{1, 2, 3},
		},
		{
			args: args{
				array: []int{1, 2, 3},
				n:     0,
			},
			want: []int{},
		},
		{
			args: args{
				array: []int{},
				n:     1,
			},
			want: []int{},
		},
		{
			args: args{
				array: []int{1, 2, 3},
				n:     0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Take(tt.args.array, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Take() = %v, want %v", got, tt.want)
			}
		})
	}
}
