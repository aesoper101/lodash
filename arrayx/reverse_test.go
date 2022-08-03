package arrayx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverse(t *testing.T) {
	result := Reverse([]int{1, 2, 3})
	require.Equal(t, []int{3, 2, 1}, result)
}
