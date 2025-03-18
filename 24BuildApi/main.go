package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")

	r := mux.NewRouter()

	//seeding

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{Fullname: "Pratyush", Website: "abc.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "MERN", CoursePrice: 199, Author: &Author{Fullname: "Pratyush", Website: "xyz.com"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers

// Serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my first Go lang API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// Grab ID from request
	params := mux.Vars(r)
	fmt.Println(params["id"])

	// loop through the courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No course Found with given Id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some  data")
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data is inside JSON")
		return
	}

	// TODO : check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// Generate unique id, convert into string
	// Append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response when id is not found
	json.NewEncoder(w).Encode("No course found with the given id")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove(index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
