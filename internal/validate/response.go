package validate

func ResponseReturn[T any](
	err error,
	obj T,
) map[string]any {

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	return map[string]any{
		"error":  errMsg,
		"result": obj,
	}
}
