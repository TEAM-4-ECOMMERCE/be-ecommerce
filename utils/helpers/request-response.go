package helpers

func SuccessGetResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": msg,
		"data":    data,
	}
}

func SuccessGetResponseData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": msg,
	}
}

func FailedResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"message": message,
		"data":    nil,
	}
}
