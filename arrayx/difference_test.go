package arrayx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testdata struct {
	Name string
}

type testdata2 struct {
	x string
	y string
}

func TestContains(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data2 := 1
	require.Equal(t, true, Contains(data1, data2))

	data3 := []testdata{
		{Name: "a"},
		{Name: "b"},
	}
	data4 := testdata{Name: "a"}
	data5 := testdata{Name: "c"}
	require.Equal(t, true, Contains(data3, data4))
	require.Equal(t, false, Contains(data3, data5))

	data6 := []*testdata{
		{Name: "a"},
		{Name: "b"},
	}
	data7 := &testdata{Name: "a"}
	data8 := &testdata{Name: "c"}
	require.Equal(t, true, Contains(data6, data7))
	require.Equal(t, false, Contains(data6, data8))
}

func TestDifference(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	require.Equal(t, []int{9, 10}, Difference(data1, data2))
}

func TestDifferenceBy(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	require.Equal(t, []int{9, 10}, DifferenceBy(data1, data2, func(i int) any {
		return i
	}))

	data3 := []testdata{
		{Name: "a"},
		{Name: "b"},
	}
	data4 := []testdata{
		{Name: "a"},
		{Name: "b"},
	}
	require.Equal(t, []testdata{}, DifferenceBy(data3, data4, func(i testdata) any {
		return i
	}))
}

func TestDifferenceWith(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	require.Equal(t, []int{9, 10}, DifferenceWith(data1, data2, func(i int, j int) bool {
		return i == j
	}))

	data3 := []testdata2{
		{x: "a", y: "b"},
	}
	data4 := []testdata2{
		{x: "b", y: "a"},
	}
	require.Equal(t, []testdata2{
		{x: "a", y: "b"},
	}, DifferenceWith(data3, data4, func(i testdata2, j testdata2) bool {
		return i.x == j.x && i.y == j.y
	}))
}

func TestContainsBy(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data2 := 1
	require.Equal(t, true, ContainsBy(data1, data2, func(i int) any {
		return i
	}))

	data3 := []testdata{
		{Name: "a"},
		{Name: "b"},
	}
	data4 := testdata{Name: "a"}
	data5 := testdata{Name: "c"}
	require.Equal(t, true, ContainsBy(data3, data4, func(i testdata) any {
		return i
	}))
	require.Equal(t, false, ContainsBy(data3, data5, func(i testdata) any {
		return i
	}))

	data6 := []*testdata{
		{Name: "a"},
		{Name: "b"},
	}
	data7 := &testdata{Name: "a"}
	data8 := &testdata{Name: "c"}
	require.Equal(t, true, ContainsBy(data6, data7, func(i *testdata) any {
		return i
	}))
	require.Equal(t, false, ContainsBy(data6, data8, func(i *testdata) any {
		return i
	}))

}

func TestContainsWith(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data2 := 1
	require.Equal(t, true, ContainsWith(data1, data2, func(i int, j int) bool {
		return i == j
	}))

	data3 := []testdata2{
		{x: "a", y: "b"},
	}
	data4 := testdata2{
		x: "a",
		y: "b",
	}
	require.Equal(t, true, ContainsWith(data3, data4, func(i testdata2, j testdata2) bool {
		return i.x == j.x && i.y == j.y
	}))
}
