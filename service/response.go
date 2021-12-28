package service

type Response struct {
	error Error
}

func (response *Response) Error(stage Stage, message string) *Response {
	response.error = *newError(stage, message)
	return response
}
