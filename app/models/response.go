package models

type (
	Response struct {
		Status		int 		`json:"status"`
		Message		string		`json:"message"`
		Content		interface{}	`json:"content"`
	}

	PaginatedResponse struct {
		Count		int			`json:"count"`
		Next		string		`json:"next"`
		Previous	string		`json:"previous"`
		Results		interface{}	`json:"results"`	
	}
)

func NewResponse(status int, message string, content interface{}) *Response {
	return &Response{
		Status: status,
		Message: message,
		Content: content,
	}
}

func NewPaginatedResponse(status, count int, message, next, prev string, results interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Content: &PaginatedResponse{
			Count:    count,
			Next:     next,
			Previous: prev,
			Results:  results,
		},
	}
}