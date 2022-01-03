package models

import (
	"time"
)

type Course struct {
	Name 		string		`json:"name" bson:"name" validate:"required"`
	Code		string		`json:"code" bson:"code" validate:"required"`
	Credits		int 		`json:"course_credits" bson:"course_credits" validate:"required,gt=0"`
	Grade		string		`json:"grade" bson:"grade" validate:"required"`
}

type Semester struct {
	Number		int			`json:"number" bson:"number" validate:"required"`
	Credits		int			`json:"earned_credits" bson:"earned_credits" validate:"required,gt=0"`
	SGPA		float32		`json:"sgpa" bson:"sgpa" validate:"required"`
	CGPA		float32		`json:"cgpa" bson:"cgpa" validate:"required"`
	Courses		[]Course	`json:"courses" bson:"courses" validatre:"required,dive"`
}

type Student struct {
	Roll		string		`json:"roll_no" bson:"roll_no" validate:"required,numeric,len=9"`
	Name		string		`json:"name" bson:"name" validate:"required"`
	Program		string		`json:"program" bson:"program" validate:"required"`
	Branch		string		`json:"branch" bson:"branch" validate:"required"`
	CGPA		float32		`json:"cgpa" bson:"cgpa" validate:"required"`
	Semesters 	[]Semester	`json:"semesters" bson:"semesters" validate:"required,dive"`
	CreatedAt	time.Time	`json:"created_at" bson:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at" bson:"updated_at"`
}