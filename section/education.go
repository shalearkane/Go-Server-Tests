package section

import (
	"bytes"
	"text/template"
	"time"
)

type Education struct {
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Institution string    `json:"institution"`
	Location    string    `json:"Location"`
	Course      string    `json:"course"`
	Degree      string    `json:"degree"`
}

func prepareEducationSection(e Education) string {
	parsedTemplate, _ := template.ParseFiles("tex/education.tex")
	buf := new(bytes.Buffer)
	parsedTemplate.Execute(buf, e)
	return buf.String()
}
