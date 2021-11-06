package models

type Response struct {
	Status		int			`json:"status"`
	Message		string		`json:"message"`
	Content		interface{}	`json:"content"`
}

func NewResponse(status int, message string, content interface{}) *Response {
	return &Response{
		Status: status,
		Message: message,
		Content: content,
	}
}