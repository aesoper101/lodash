package utils

func Over[T any](predicates ...func(T) T) func(T) []T {
	return func(value T) []T {
		results := make([]T, 0)
		for _, predicate := range predicates {
			results = append(results, predicate(value))
		}
		return results
	}
}

func OverEvery[T any](predicates ...func(T) bool) func(T) bool {
	return func(value T) bool {
		for _, predicate := range predicates {
			if !predicate(value) {
				return false
			}
		}
		return true
	}
}

func OverSome[T any](predicates ...func(T) bool) func(T) bool {
	return func(value T) bool {
		for _, predicate := range predicates {
			if predicate(value) {
				return true
			}
		}
		return false
	}
}
