package arrayx

// Chunk  Creates an arrayx of elements split into groups the length of size.
// if arrayx can't be split evenly, the final chunk will be the remaining elements.
func Chunk[T any](array []T, size int) [][]T {
	chunks := make([][]T, 0)

	for i := 0; i < len(array); i++ {
		index := i / size

		if i%size == 0 {
			chunks = append(chunks, make([]T, 0))
		}

		chunks[index] = append(chunks[index], array[i])
	}

	return chunks
}
