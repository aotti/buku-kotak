package helper

type THandlerResponse struct {
	Status  int `json:"status"`
	Message any `json:"message"`
	Data    any `json:"data"`
}

func HandlerResponse(status int, data any, err error) THandlerResponse {
	message := "success"
	if err != nil {
		message = err.Error()
	}

	return THandlerResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
