package utils

func Map[T1, T2 any](in []T1, f func(T1) T2) []T2 {
	res := make([]T2, len(in))
	for i, item := range in {
		res[i] = f(item)
	}
	return res
}
