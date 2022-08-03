package arrayx

import (
	"reflect"
	"testing"
)

func TestInterSection(t *testing.T) {
	//d1 := []int{1, 2, 3, 4}
	//d := DropWhile(d1, func(v int, n uint) bool {
	//	fmt.Println(v)
	//	return v < 3
	//})
	//t.Logf("%v", d)

	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				arrays: [][]int{
					{1, 2, 3},
					{1, 2, 3},
					{1, 2},
				},
			},
			want: []int{1, 2},
		},
		{
			args: args{
				arrays: [][]int{
					{2, 1},
					{2, 3},
				},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterSection(tt.args.arrays...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
