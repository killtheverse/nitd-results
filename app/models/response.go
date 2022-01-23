package models

// Response for returning a single student
// swagger:response Response
type Response struct {
	// Status of the response
	Status		int 		`json:"status"`
	// Message in the response
	Message		string		`json:"message"`
	// Student to be displayed
	Student		Student		`json:"student"`
}

// Response for returning an error message
// swagger:response ErrorResponse
type ErrorResponse struct {
	// Status of the response
	Status		int			`json:"status"`
	// Error message
	Message		string		`json:"message"`
	// Content to displayed along with the message
	Content		interface{}	`json:"content"`
}

// Response for returning a list of students
// swagger:response PaginatedResponse
type PaginatedResponse struct {
	// Status of the response
	Status		int			`json:"status"`
	// Message in the response
	Message		string		`json:"message"`
	// Count of students in response
	// maximum: 100
	// minimum: 0
	Count		int			`json:"count"`
	// Path to fetch the students next in the list
	Next		string		`json:"next"`
	// Path to fetch the students before the items currently in the list
	Previous	string		`json:"previous"`
	// List of students
	// collection format: Student
	Students	[]Student	`json:"results"`	
}

// NewResponse will return new instance of Response
func NewResponse(status int, message string, student Student) *Response {
	return &Response{
		Status: status,
		Message: message,
		Student: student,
	}
}

// NewErrorResponse will return new instance of ErrorResponse
func NewErrorResponse(status int, message string, content interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status: status,
		Message: message,
		Content: content,
	}
}

// NewPaginatedResponse will return new instance of PaginatedResponse
func NewPaginatedResponse(status, count int, message, next, prev string, students []Student) *PaginatedResponse {
	return &PaginatedResponse{
		Status:  status,
		Message: message,
		Count:    count,
		Next:     next,
		Previous: prev,
		Students:  students,
	}
}