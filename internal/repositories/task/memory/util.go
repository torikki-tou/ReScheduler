package memory

func updateField[T any](field any, def T) T {
	if field == nil {
		return def
	}

	return field.(T)
}
