package resources

func Resources[M ~[]E, D ~[]R, E, R any](models M, dest *D, operation func(E) R) {
	for _, model := range models {
		*dest = append(*dest, operation(model))
	}
}
