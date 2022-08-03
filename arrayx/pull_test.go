package arrayx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPull(t *testing.T) {
	result := Pull([]string{"a", "b", "c", "a", "b", "c"}, "a", "c")
	require.Equal(t, []string{"b", "b"}, result)

	result1 := Pull([]testdata{{Name: "1"}, {Name: "2"}, {Name: "3"}, {Name: "4"}}, testdata{Name: "3"})
	require.Equal(t, []testdata{{Name: "1"}, {Name: "2"}, {Name: "4"}}, result1)

	data := []testdata{{Name: "1"}, {Name: "2"}, {Name: "3"}, {Name: "4"}}
	result2 := Pull(data, testdata{Name: "3"})
	require.Equal(t, []testdata{{Name: "1"}, {Name: "2"}, {Name: "4"}}, result2)
	require.Equal(t, 4, len(data))

	data1 := []*testdata{{Name: "1"}, {Name: "2"}, {Name: "3"}, {Name: "4"}}
	result3 := Pull(data1, &testdata{Name: "3"})
	require.Equal(t, []*testdata{{Name: "1"}, {Name: "2"}, {Name: "4"}}, result3)
	require.Equal(t, 4, len(data1))
}

func TestPullAll(t *testing.T) {
	result := PullAll([]int{1, 2, 3, 1, 2, 3}, []int{2, 3})
	require.Equal(t, []int{1, 1}, result)
}

func TestPullAllBy(t *testing.T) {
	result := PullAllBy([]int{1, 2, 3, 1, 2, 3}, []int{2, 3}, func(a int) any {
		return a
	})
	require.Equal(t, []int{1, 1}, result)
}

func TestPullAllWith(t *testing.T) {
	result := PullAllWith([]int{1, 2, 3, 1, 2, 3}, []int{2, 3}, func(a, b int) bool {
		return a == b
	})
	require.Equal(t, []int{1, 1}, result)
}

func TestPullAt(t *testing.T) {
	result := PullAt([]string{"a", "b", "c", "d"}, 1, 3)
	require.Equal(t, []string{"b", "d"}, result)
}
