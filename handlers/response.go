package handlers

type Response struct {
	Code int
	Body Body
}

type Body struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    any               `json:"data"`
	Errors  map[string]string `json:"errors"`
}

func (Response) Success(message string) Response {
	return Response{
		Code: 200,
		Body: Body{
			Status:  "success",
			Message: message,
		},
	}
}

func (Response) SuccessWithData(message string, data any) Response {
	return Response{
		Code: 200,
		Body: Body{
			Status:  "success",
			Message: message,
			Data:    data,
		},
	}
}

func (Response) Failure(code int, message string) Response {
	return Response{
		Code: code,
		Body: Body{
			Status:  "failure",
			Message: message,
		},
	}
}

func (Response) FailureWithErrors(code int, message string, errors map[string]string) Response {
	return Response{
		Code: code,
		Body: Body{
			Status:  "failure",
			Message: message,
			Errors:  errors,
		},
	}
}

func (Response) NotFound(message string) Response {
	if message == "" {
		message = "Not found."
	}
	return Response{}.Failure(404, message)
}

func (Response) ValidationFailed(message string, errors map[string]string) Response {
	if message == "" {
		message = "Validation failed."
	}
	return Response{}.FailureWithErrors(422, message, errors)
}
