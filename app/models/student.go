package models

import (
	"encoding/json"
	"io"
	"time"
)

type Course struct {
	Name 		string		`json:"name"`
	Code		string		`json:"code"`
	Credits		int 		`json:"course_credits"`
	Grade		string		`json:"grade"`
}

type Semester struct {
	Number		int			`json:"number"`
	Credits		int			`json:"earned_credits"`
	SGPA		float32		`json:"sgpa"`
	CGPA		float32		`json:"cgpa"`
	Courses		[]Course	`json:"courses"`
}

type Student struct {
	Roll		string		`json:"roll_no"`
	Name		string		`json:"name"`
	Program		string		`json:"program"`
	Branch		string		`json:"branch"`
	CGPA		float32		`json:"cgpa"`
	Semesters 	[]Semester	`json:"semesters"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
}

func (s *Student) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(s)
}

func NewCourse(name string, code string, credits int, grade string) *Course {
	return &Course{
		Name: name,
		Code: code,
		Credits: credits,
		Grade: grade,
	}
}

func NewSemester(number int, credits int, sgpa float32, cgpa float32, courses []Course) *Semester {
	return &Semester{
		Number: number,
		Credits: credits,
		SGPA: sgpa,
		CGPA: cgpa,
		Courses: courses,
	}
}

func NewStudent(roll string, name string, program string, branch string, cgpa float32, semesters []Semester) *Student {
	return &Student{
		Name: name,
		Roll: roll,
		Program: program,
		Branch: branch,
		CGPA: cgpa,
		Semesters: semesters,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}