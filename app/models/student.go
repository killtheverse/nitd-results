package models

import (
	"encoding/json"
	"io"
	"time"
)

type Student struct {
	Roll		string		`json:"roll"`
	Name		string		`json:"name"`
	Branch		string		`json:"branch"`
	Semester	int			`json:"semester"`
	Result 		[]float32	`json:"result"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
}

func (s *Student) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(s)
}

