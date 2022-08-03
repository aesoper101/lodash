package mathx

func Max[T Number](numbers ...T) T {
	if len(numbers) == 0 {
		panic("min requires at least one number")
	}
	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}
