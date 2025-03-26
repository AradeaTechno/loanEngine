package helpers

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ApiResponse(err error, data any) Response {
	var message string
	if err != nil {
		message = err.Error()
	} else {
		message = "request success"
	}
	if data == nil {
		return Response{
			Message: message,
		}
	}
	return Response{
		Message: message,
		Data:    data,
	}
}
