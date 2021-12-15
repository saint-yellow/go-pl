package main

import (
	htmlTemplate "html/template"
	"isbn-9787111558422/src/chapter04/github"
	"log"
	"os"
)

const template = `
<h1>{{.TotalCount}} issues</h1>
<table>
    <tr style="text-align: left">
        <th>#</th>
        <th>State</th>
        <th>User</th>
        <th>Title</th>
    </tr>
    {{range .Items}}
    <tr>
        <td><a href="{{.HTMLURL}}">{{.Number}}</a></td>
        <td>{{.State}}</td>
        <td><a href="{{.User.HTMLURL}}">{{.User.Login}}</a></td>
        <td><a href="{{.HTMLURL}}">{{.Title}}</a></td>
    </tr>
    {{end}}
</table>
`

func main() {
	var issueList = htmlTemplate.Must(htmlTemplate.New("issuelist").Parse(template))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
