package arrayx

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should return the union of the given arrays",
			args: args{
				arrays: [][]int{
					{1, 2, 3},
					{2, 3, 4},
					{3, 4, 5},
				},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			args: args{
				arrays: [][]int{
					{2},
					{1, 2},
				},
			},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.arrays...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}
