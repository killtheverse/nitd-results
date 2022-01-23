package models

import (
	"time"
)

// Course defines the structure of a course
// swagger:model
type Course struct {
	// name of the course
	// required: true
	Name 		string		`json:"name" bson:"name" validate:"required"`
	
	// code for the course
	// required: true
	Code		string		`json:"code" bson:"code" validate:"required"`
	
	// number of credits for the course
	// required: true
	// min: 1
	Credits		int 		`json:"course_credits" bson:"course_credits" validate:"required,gt=0"`
	
	// grade received by the student
	// required: true
	Grade		string		`json:"grade" bson:"grade" validate:"required"`
}

// Semester defines the structure of a semester
// swagger:model
type Semester struct {
	// semester number
	// required: true
	Number		int			`json:"number" bson:"number" validate:"required"`
	
	// credits earned by the student in that semester
	// required: true
	// min: 1
	Credits		int			`json:"earned_credits" bson:"earned_credits" validate:"required,gt=0"`
	
	// sgpa earned for the semester by the student
	// required: true
	SGPA		float32		`json:"sgpa" bson:"sgpa" validate:"required"`
	
	// cgpa after this semester
	// required: true
	CGPA		float32		`json:"cgpa" bson:"cgpa" validate:"required"`
	
	// information about courses taken by the student in the semester
	// required: true
	Courses		[]Course	`json:"courses" bson:"courses" validatre:"required,dive"`
}

// Student defines the structure of a student
// swagger:model
type Student struct {
	// roll number of student
	// required: true
	Roll		string		`json:"roll_no" bson:"roll_no" validate:"required,numeric,len=9"`
	
	// name of the student
	// required: true
	Name		string		`json:"name" bson:"name" validate:"required"`
	
	// program in which the student is enrolled in 
	// required: true
	Program		string		`json:"program" bson:"program" validate:"required"`
	
	// branch of the student
	// required: true
	Branch		string		`json:"branch" bson:"branch" validate:"required"`
	
	// current cgpa of the student
	// required: true
	CGPA		float32		`json:"cgpa" bson:"cgpa" validate:"required"`
	
	// information about the previous semesters of a student
	// required: true
	Semesters 	[]Semester	`json:"semesters" bson:"semesters" validate:"required,dive"`
	
	// time at which the instance was created in the database
	// required: false
	CreatedAt	time.Time	`json:"created_at" bson:"created_at"`
	
	// time at which the student instance was last updated in the database
	// required: false
	UpdatedAt	time.Time	`json:"updated_at" bson:"updated_at"`
}

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

// swagger:parameters studentDetail
type studentDetailParameterWrapper struct {
	// The roll number of the student
	// in: path
	// required: true
	Roll	string	`json:"roll_number"`
}