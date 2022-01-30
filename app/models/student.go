package models

import (
	"time"
)

// Course defines the structure of a course
// swagger:model
type Course struct {
	// Name of the course
	// required: true
	Name 		string		`json:"name" bson:"name" validate:"required"`
	
	// Code for the course
	// required: true
	Code		string		`json:"code" bson:"code" validate:"required"`
	
	// Number of credits for the course
	// required: true
	// min: 1
	Credits		int 		`json:"course_credits" bson:"course_credits" validate:"required,gt=0"`
	
	// Grade received by the student
	// required: true
	Grade		string		`json:"grade" bson:"grade" validate:"required"`
}

// Semester defines the structure of a semester
// swagger:model
type Semester struct {
	// Semester number
	// required: true
	Number		int			`json:"number" bson:"number" validate:"required"`
	
	// Credits earned by the student in that semester
	// required: true
	// min: 1
	Credits		int			`json:"earned_credits" bson:"earned_credits" validate:"required,gt=0"`
	
	// SGPA earned for the semester by the student
	// required: true
	SGPA		float32		`json:"sgpa" bson:"sgpa" validate:"required"`
	
	// CGPA after this semester
	// required: true
	CGPA		float32		`json:"cgpa" bson:"cgpa" validate:"required"`
	
	// Information about courses taken by the student in the semester
	// required: true
	Courses		[]Course	`json:"courses" bson:"courses" validatre:"required,dive"`
}

// Student defines the structure of a student
// swagger:model
type Student struct {
	// Roll number of student
	// required: true
	Roll		string		`json:"roll_no" bson:"roll_no" validate:"required,numeric,len=9"`
	
	// Name of the student
	// required: true
	Name		string		`json:"name" bson:"name" validate:"required"`
	
	// Program in which the student is enrolled in 
	// required: true
	Program		string		`json:"program" bson:"program" validate:"required"`
	
	// Branch of the student
	// required: true
	Branch		string		`json:"branch" bson:"branch" validate:"required"`
	
	// Current CGPA of the student
	// required: true
	CGPA		float32		`json:"cgpa" bson:"cgpa" validate:"required"`
	
	// Information about the previous semesters of a student
	// required: true
	Semesters 	[]Semester	`json:"semesters" bson:"semesters" validate:"required,dive"`
	
	// Time at which the instance was created in the database
	// required: false
	// read only: true
	CreatedAt	time.Time	`json:"created_at" bson:"created_at"`
	
	// Time at which the student instance was last updated in the database
	// required: false
	// read only: true
	UpdatedAt	time.Time	`json:"updated_at" bson:"updated_at"`
}
