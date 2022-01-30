// Package classification for NITD Results API
//
// Documentation for NITD Results API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//	Contact: Rahul Dev Kureel<r.dev2000@gmail.com>
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"github.com/killtheverse/nitd-results/app/models"
)

// swagger:parameters listStudents
type studentListParameterWrapper struct {
	// The branch of the students
	// in: query
	// required: false
	Branch	string	`json:"branch"`

	// The program of the students
	// in: query
	// required: false
	Program	string	`json:"program"`

	// Limit on the number of results
	// in: query
	// required: false
	Limit	int		`json:"limit"`

	// Offset on the results
	// in: query
	// required: false
	Offset	int		`json:"offset"`	
}

// swagger:parameters studentDetail updateStudent
type studentRollParameterWrapper struct {
	// The roll number of the student
	// in: path
	// required: true
	Roll	string	`json:"roll_number"`
}

// swagger:response Response
type responseWrapper struct {
	// Response for returning a single student
	// in: body
	Body models.Response
}

// swagger:response ErrorResponse
type errorResponseWrapper struct {
	// Response for returning an error
	// in: body
	Body models.ErrorResponse
}

// swagger:response PaginatedResponse
type paginatedResponseWrapper struct {
	// Response for returning a list of students with pagination
	// in: body
	Body models.PaginatedResponse
}