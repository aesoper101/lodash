package arrayx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDrop(t *testing.T) {
	data := []int{1, 2, 3, 4}
	result := Drop(data, 2)
	require.Equal(t, []int{3, 4}, result)

	result = Drop(data, 0)
	require.Equal(t, data, result)

	result = Drop(data, 5)
	require.Equal(t, []int{}, result)

	result = Drop(data)
	require.Equal(t, []int{2, 3, 4}, result)
}

func TestDropRight(t *testing.T) {
	data := []int{1, 2, 3, 4}
	result := DropRight(data, 2)
	require.Equal(t, []int{1, 2}, result)

	result = DropRight(data, 0)
	require.Equal(t, data, result)

	result = DropRight(data, 5)
	require.Equal(t, []int{}, result)

	result = DropRight(data)
	require.Equal(t, []int{1, 2, 3}, result)
}

type user struct {
	user   string
	active bool
}

func TestDropRightWhile(t *testing.T) {
	var users = []user{
		{user: "barney", active: true},
		{user: "fred", active: false},
		{user: "pebbles", active: false},
	}

	result := DropRightWhile(users, func(u user, n uint) bool {
		return !u.active
	})
	require.Equal(t, []user{
		{user: "barney", active: true},
	}, result)

	result = DropRightWhile(users, func(u user, n uint) bool {
		return u.user == "pebbles" && u.active == false
	})
	require.Equal(t, []user{
		{user: "barney", active: true},
		{user: "fred", active: false},
	}, result)

	result = DropRightWhile(users, func(u user, n uint) bool {
		return !u.active
	})
	require.Equal(t, []user{
		{user: "barney", active: true},
	}, result)

	result = DropRightWhile(users, func(u user, n uint) bool {
		return u.active
	})
	require.Equal(t, users, result)
}

func TestDropWhile(t *testing.T) {
	var users = []user{
		{user: "barney", active: false},
		{user: "fred", active: false},
		{user: "pebbles", active: true},
	}

	result := DropWhile(users, func(u user, n uint) bool {
		return !u.active
	})
	require.Equal(t, []user{
		{user: "pebbles", active: true},
	}, result)

	result = DropWhile(users, func(u user, n uint) bool {
		return u.user == "barney" && u.active == false
	})
	require.Equal(t, []user{
		{user: "fred", active: false},
		{user: "pebbles", active: true},
	}, result)

	result = DropWhile(users, func(u user, n uint) bool {
		return u.active
	})
	require.Equal(t, users, result)
}
