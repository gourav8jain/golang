package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

// fake db
var newcourse []NewCourse

func (c *NewCourse) IsEmpty() bool {
	return c.CourseId == "0" && c.Name == ""
}

func main() {
	url := "https://lco.dev"
	fmt.Println("Website")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response after read", string(databytes))

	// server started on 8000 and call get calls and others -- GET

	fmt.Println("Welcome to GET")
	GetRequest()

	// server started on 8000 and call get calls and others -- POST

	fmt.Println("Welcome to POST")
	PostRequest()

	// server started on 8000 and call get calls and others -- POST

	fmt.Println("Welcome to POST FORM")
	PostFormRequest()

	// ENCODE JSON DATA

	fmt.Println("Welcome to ENCODE JSON DATA")
	EncodeJson()

	// ENCODE JSON DATA

	fmt.Println("Welcome to DECODE JSON DATA")
	DecodeJson()

	// MOD - module to create an API
	fmt.Println("Mod users")
	// r := mux.NewRouter()
	// r.HandleFunc("/", ServeHome).Methods("GET")

	// log.Fatal(http.ListenAndServe(":4000", r))

	// seeding

	newcourse = append(newcourse, NewCourse{CourseId: "1", Name: "course-1", Price: 22,
		Author: &author{Name: "Author-1", AuthorId: "1"}})
	newcourse = append(newcourse, NewCourse{CourseId: "2", Name: "course-2", Price: 22,
		Author: &author{Name: "Author-2", AuthorId: "2"}})

	// build API
	fmt.Println("Proper API")

	// routing
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome1).Methods("GET")
	router.HandleFunc("/courses", getallcourses).Methods("GET")
	router.HandleFunc("/course/{id}", getonecourse).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", router))

}

func GetRequest() {
	const url = "http://localhost:8000/"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code", response.StatusCode)
	fmt.Println("Content length", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Content", string(content))

	defer response.Body.Close()
}

func PostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	data := url.Values{}

	data.Add("name", "Gourav")
	data.Add("age", "10")

	fmt.Println("Data", data)
	responsebody, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code", responsebody.StatusCode)
	fmt.Println("Content length", responsebody.ContentLength)

	content, err := ioutil.ReadAll(responsebody.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Content", string(content))

	defer responsebody.Body.Close()

}

func PostRequest() {
	const url = "http://localhost:8000/post"
	body := strings.NewReader(`
	{
		"name":"Gourav",
		"age":10
		}`)

	fmt.Println("Body", body)
	responsebody, err := http.Post(url, "application/json", body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code", responsebody.StatusCode)
	fmt.Println("Content length", responsebody.ContentLength)

	content, err := ioutil.ReadAll(responsebody.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Content", string(content))

	defer responsebody.Body.Close()

}

func EncodeJson() {
	courseObj := []course{
		{1, "Course-1", 10, []string{"tag-1", "tag-2"}},
		{2, "Course-2", 10, []string{"tag-1", "tag-2"}},
		{3, "Course-3", 10, nil},
	}

	//finalJson, err := json.Marshal(courseObj)
	finalJson, err := json.MarshalIndent(courseObj, "", "\t")
	if err != nil {
		panic(nil)
	}
	fmt.Println("Final JSON - Address", finalJson)
	fmt.Println("Final JSON", string(finalJson))
	fmt.Printf("Final JSON %s", finalJson)
}

func DecodeJson() {
	jsondatafromweb := []byte(`
		{
			"course-id": 3,
			"Name": "Course-3",
			"Price": 10,
			"Tags": ["tag-1","tag-2"]
		}
	`)

	var lcocourse course

	jsonIsValid := json.Valid(jsondatafromweb)
	fmt.Println("IsValid", jsonIsValid)

	if jsonIsValid {
		json.Unmarshal(jsondatafromweb, &lcocourse)
		fmt.Println("Unmarshal", lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	}
}

func greeter() {
	fmt.Println("Mod users")
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome</h1>"))
}

type course struct {
	Id    int `json:"course-id"`
	Name  string
	Price int
	Tags  []string
}

type NewCourse struct {
	CourseId string `json:"newcourse-id"`
	Name     string
	Price    int
	Author   *author
}

type author struct {
	Name     string
	AuthorId string `json:"author-id"`
}

// controller - go into seperate files
// serve home route

func serveHome1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to mod api</h1>"))
	w.Header().Set("Content-Type", "application/json")
}

func getallcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newcourse)
}

func getonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	fmt.Println("Params", params)

	for _, course := range newcourse {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	fmt.Println("No course found for particular course id")
}

func createonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	var newcourse NewCourse
	_ = json.NewDecoder(r.Body).Decode(&newcourse)
}
