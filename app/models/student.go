package models

import (
	"time"
)

type Course struct {
	Name 		string		`json:"name" validate:"required"`
	Code		string		`json:"code" validate:"required"`
	Credits		int 		`json:"course_credits" validate:"required,gt=0"`
	Grade		string		`json:"grade" validate:"required"`
}

type Semester struct {
	Number		int			`json:"number" validate:"required"`
	Credits		int			`json:"earned_credits" validate:"required,gt=0"`
	SGPA		float32		`json:"sgpa" validate:"required"`
	CGPA		float32		`json:"cgpa" validate:"required"`
	Courses		[]Course	`json:"courses" validatre:"required,dive"`
}

type Student struct {
	Roll		string		`json:"roll_no" validate:"required,numeric,len=9"`
	Name		string		`json:"name" validate:"required"`
	Program		string		`json:"program" validate:"required"`
	Branch		string		`json:"branch" validate:"required"`
	CGPA		float32		`json:"cgpa" validate:"required"`
	Semesters 	[]Semester	`json:"semesters" validate:"required,dive"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
}