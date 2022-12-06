package middleware

type Response struct {
	Data  interface{}    `json:"data"`
	Error *ResponseError `json:"error"`
}

type ResponseError struct {
	Message string `json:"message,omitempty"`
}

func NewResponse(data interface{}) *Response {
	return &Response{
		Data: data,
	}
}

func NewErrorResponse(err error) *Response {
	return &Response{
		Error: &ResponseError{
			Message: err.Error(),
		},
	}
}
