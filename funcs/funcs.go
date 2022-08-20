package funcs

func Map[T any, V any](slice []T, mapFunc func(T) V) []V {
	var newSlice []V
	for _, el := range slice {
		newEl := mapFunc(el)
		newSlice = append(newSlice, newEl)
	}
	return newSlice
}

func Filter[T any](slice[]T, filterFunc func(T) bool) []T {
	var newSlice []T
	for _, el := range slice {
		if filterFunc(el) {
			newSlice = append(newSlice, el)
		}
	}
	return newSlice
}

func Reduce[T any, V any](slice[]T, reducerFunc func(V, T) V, initValue V) V {
	for _, el := range slice {
		initValue = reducerFunc(initValue, el)
	}
	return initValue
}