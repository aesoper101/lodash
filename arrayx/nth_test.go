package arrayx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNth(t *testing.T) {
	result := Nth[string]([]string{"a", "b", "c", "d"}, 2)
	require.Equal(t, "c", result)

	result = Nth[string]([]string{"a", "b", "c", "d"}, 1)
	require.Equal(t, "b", result)

	result = Nth[string]([]string{"a", "b", "c", "d"}, -2)
	require.Equal(t, "c", result)

	result = Nth[string]([]string{"a", "b", "c", "d"}, -5)
	require.Equal(t, "", result)

	result2 := Nth([]testdata{{Name: "1"}, {Name: "2"}, {Name: "3"}, {Name: "4"}}, -1)
	require.Equal(t, "4", result2.Name)

	result3 := Nth([]testdata{{Name: "1"}, {Name: "2"}, {Name: "3"}, {Name: "4"}}, 0)
	require.Equal(t, "1", result3.Name)

	result4 := Nth([]testdata{{Name: "1"}, {Name: "2"}, {Name: "3"}, {Name: "4"}}, 1)
	require.Equal(t, "2", result4.Name)

	result5 := Nth([]testdata{{Name: "1"}, {Name: "2"}, {Name: "3"}, {Name: "4"}}, 5)
	require.Equal(t, testdata{}, result5)
	require.Equal(t, "", result5.Name)
}
