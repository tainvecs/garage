package middleware

// Response is the struct of API response
type Response struct {
	Data  interface{}    `json:"data"`
	Error *ResponseError `json:"error"`
}

// ResponseError is part of the Response
type ResponseError struct {
	Message string `json:"message,omitempty"`
}

// NewResponse creates a new Response
func NewResponse(data interface{}) *Response {
	return &Response{
		Data: data,
	}
}

// NewErrorResponse creates a new Response when error happens
func NewErrorResponse(err error) *Response {
	return &Response{
		Error: &ResponseError{
			Message: err.Error(),
		},
	}
}
