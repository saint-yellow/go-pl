package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/saint-yellow/go-pl/toys/distributed-system/business"
	"github.com/saint-yellow/go-pl/toys/distributed-system/registry"
)

func RegisterHandlers() {
	http.Handle("/", http.RedirectHandler("/students", http.StatusPermanentRedirect))

	h := new(studentsHandler)
	http.Handle("/students", h)
	http.Handle("/students/", h)
}

type studentsHandler struct{

}

func (sh studentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	switch len(pathSegments) {
	case 2:
		sh.renderStudents(w, r)
	case 3:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.renderStudent(w, r, id)
	case 4:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if strings.ToLower(pathSegments[3]) != "grades" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.renderGrades(w, r, id)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (sh studentsHandler) renderStudents(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("error retrieving students: ", err)
		}
	}()

	serviceURL, err := registry.GetProvider(registry.BusinessService)
	if err != nil {
		return
	}

	res, err := http.Get(serviceURL + "/students")
	if err != nil {
		return
	}

	var s business.Students
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return
	}

	rootTemplate.Lookup("students.html").Execute(w, s)
}

func (sh studentsHandler) renderStudent(w http.ResponseWriter, r *http.Request, id int) {
	var err error

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("error retrieving student: ", err)
		}
	}()

	serviceURL, err := registry.GetProvider(registry.BusinessService)
	if err != nil {
		return
	}

	res, err := http.Get(fmt.Sprintf("%v/students/%v", serviceURL, id))
	if err != nil {
		return
	}

	var s business.Student
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return
	}

	rootTemplate.Lookup("student.html").Execute(w, s)
}

func (sh studentsHandler) renderGrades(w http.ResponseWriter, r *http.Request, id int) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	defer func() {
		w.Header().Add("location", fmt.Sprintf("/students/%v", id))
		w.WriteHeader(http.StatusTemporaryRedirect)
	}()

	title := r.FormValue("Title")
	gradeType := r.FormValue("Type")
	score, err := strconv.ParseFloat(r.FormValue("Score"), 32)
	if err != nil {
		log.Println("failed to parse score: ", err)
		return
	}

	g := business.Grade{
		Title: title,
		Type: business.GradeType(gradeType),
		Score: float32(score),
	}
	data, err := json.Marshal(g)
	if err != nil {
		log.Println("failed to convert grade to JSON: ", g, err)
	}

	serviceURL, err := registry.GetProvider(registry.BusinessService)
	if err != nil {
		log.Println("failed to retrieve instance of business service", err)
		return
	}

	res, err := http.Post(fmt.Sprintf("%v'students/%v/grades", serviceURL, id), "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("failed to save grade to business service", err)
		return
	}
	if res.StatusCode != http.StatusCreated {
		log.Println("failed to save grade to business service. status code: ", res.StatusCode)
		return
	}
}