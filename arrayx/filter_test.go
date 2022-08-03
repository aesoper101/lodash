package arrayx

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		array     []int
		predicate func(int, int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				array: []int{1, 2, 3, 4, 5},
				predicate: func(value, index int) bool {
					return value > 3
				},
			},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.array, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
