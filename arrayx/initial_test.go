package arrayx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInitial(t *testing.T) {
	list := []int{1, 2, 3}
	result := Initial(list)
	require.Equal(t, []int{1, 2}, result)

	list1 := []int{1}
	result1 := Initial(list1)
	require.Equal(t, []int{}, result1)

	list2 := make([]int, 0)
	result2 := Initial(list2)
	require.Equal(t, []int{}, result2)
}
