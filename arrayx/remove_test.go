package arrayx

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	type args struct {
		arr       []int
		predicate func(int2 int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				arr: []int{1, 2, 3, 4},
				predicate: func(int2 int) bool {
					return int2%2 == 0
				},
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.arr, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
