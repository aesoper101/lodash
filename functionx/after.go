package functionx

// After creates a functionx that invokes func once it's called n or more times.
func After(n int, fn func()) func() {
	if n <= 0 {
		return func() {}
	}
	return func() {
		if n--; n > 0 {
			return
		}
		fn()
	}
}

// Before Creates a functionx that invokes func, with the this binding and arguments of the created functionx,
// while it's called less than n times. Subsequent calls to the created functionx return the result of
// the last func invocation.
func Before(n int, fn func()) func() {
	if n <= 0 {
		return func() {}
	}
	return func() {
		if n--; n > 0 {
			return
		}
		fn()
	}
}
