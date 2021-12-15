package main

import (
	"isbn-9787111558422/src/chapter04/github"
	"log"
	"os"
	textTemplate "text/template"
	"time"
)

const template = `
{{.TotalCount}} issues:
{{range .Items}}------------------------
Number:		{{.Number}}
User:		{{.User.Login}}
Title:		{{.Title | printf "%.64s"}}
Age:		{{.CreatedAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	report, err := textTemplate.New("report").Funcs(textTemplate.FuncMap{"daysAgo": daysAgo}).Parse(template)
	if err != nil {
		log.Fatal(err)
	}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}