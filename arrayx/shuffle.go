package arrayx

import "math/rand"

func Shuffle[T any](a []T) {
	n := len(a)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
