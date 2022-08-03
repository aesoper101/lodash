package arrayx

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	type args struct {
		array []string
		size  int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			args: args{
				array: []string{"a", "b", "c", "d"},
				size:  2,
			},
			want: [][]string{
				{"a", "b"},
				{"c", "d"},
			},
		},
		{
			args: args{
				array: []string{"a", "b", "c", "d"},
				size:  3,
			},
			want: [][]string{
				{"a", "b", "c"},
				{"d"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk[string](tt.args.array, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
