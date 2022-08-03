package functionx

import "github.com/aesoper101/lodash/arrayx"

func Keys[K comparable, V any](object map[K]V) []K {
	results := make([]K, 0, len(object))

	for key := range object {
		results = append(results, key)
	}

	return results
}

func Values[K comparable, V any](object map[K]V) []V {
	results := make([]V, 0, len(object))

	for _, value := range object {
		results = append(results, value)
	}

	return results
}

func Pick[K comparable, V any](object map[K]V, keys ...K) map[K]V {
	result := make(map[K]V, len(keys))

	for _, key := range keys {
		if value, ok := object[key]; ok {
			result[key] = value
		}
	}

	return result
}

func PickBy[K comparable, V any](object map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V)

	for key, value := range object {
		if predicate(key, value) {
			result[key] = value
		}
	}

	return result
}

func Omit[K comparable, V any](object map[K]V, keys ...K) map[K]V {
	result := make(map[K]V, len(object))

	for key, value := range object {
		if !arrayx.Contains(keys, key) {
			result[key] = value
		}
	}

	return result
}

func OmitBy[K comparable, V any](object map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V)

	for key, value := range object {
		if !predicate(key, value) {
			result[key] = value
		}
	}

	return result
}

func MapKeys[K comparable, V any, R comparable](object map[K]V, iteratee func(K, V) R) map[R]V {
	result := make(map[R]V)

	for key, value := range object {
		result[iteratee(key, value)] = value
	}

	return result
}

func MapValues[K comparable, V any, R any](object map[K]V, iteratee func(K, V) R) map[K]R {
	result := make(map[K]R)

	for key, value := range object {
		result[key] = iteratee(key, value)
	}

	return result
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func Entries[K comparable, V any](object map[K]V) []Entry[K, V] {
	results := make([]Entry[K, V], 0, len(object))

	for key, value := range object {
		results = append(results, Entry[K, V]{Key: key, Value: value})
	}

	return results
}

func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	result := make(map[K]V)

	for _, entry := range entries {
		result[entry.Key] = entry.Value
	}

	return result
}

func Assign[K comparable, V any](object map[K]V, others ...map[K]V) map[K]V {
	for _, other := range others {
		for key, value := range other {
			object[key] = value
		}
	}

	return object
}
