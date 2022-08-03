package collectionx

func FindKey[V any, K comparable](object map[K]V, predicate func(V) bool) K {
	for key, value := range object {
		if predicate(value) {
			return key
		}
	}
	var nilKey K
	return nilKey
}

func FindLastKey[V any, K comparable](object map[K]V, predicate func(V) bool) K {
	var result K
	for key, value := range object {
		if predicate(value) {
			result = key
		}
	}

	return result
}
