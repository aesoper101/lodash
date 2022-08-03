package collectionx

func ForIn[T any, K comparable](object map[K]T, iteratee func(key K, value T)) {
	for key, value := range object {
		iteratee(key, value)
	}
}

func ForInRight[T any, K comparable](object map[K]T, iteratee func(value T, key K)) {
	for key := range object {
		iteratee(object[key], key)
	}
}
