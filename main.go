package main

import (
	"context"

	"time"

	"fmt"

	"strconv"

	"log"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	"html/template"

	"Personal-Web/connection"
)

var Data = map[string]interface{}{
	"Title":   "Personal Web",
	"IsLogin": false,
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type Project struct {
	Id           int
	Title        string
	Start_date   time.Time
	End_date     time.Time
	Description  string
	Technologies string
	NodeJs       string
	Java         string
	Php          string
	Laravel      string
	Image        string
	Format_start string
	Format_end   string
}

//type Projects []Project

//func NewProject() *Project {
//	return &Project{
//		Start_date: time.Date(2022, 5, 12, 21, 0, 0, 0, time.Local),
//		End_date:   time.Date(2022, 5, 12, 21, 0, 0, 0, time.Local),
//	}
//}

//var Projects = []Project{
//	{
//		Title:       "Pembelajaran Online",
//		Duration:    "Duration : 3 Weeks",
//		Author:      " | Bagas",
//		Description: "Sangat sulit sekali hehehehe",
//	},
//}

// function routing
func main() {
	router := mux.NewRouter()

	connection.DatabaseConnect()

	// Create Folder
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/Project", project).Methods("GET")
	router.HandleFunc("/addProject", addProject).Methods("POST")
	router.HandleFunc("/contactMe", contactMe).Methods("GET")
	router.HandleFunc("/projectDetail/{id}", projectDetail).Methods("GET")
	router.HandleFunc("/delete-project/{id}", deleteProject).Methods("GET")
	router.HandleFunc("/signForm", signupForm).Methods("GET")
	router.HandleFunc("/signUp", signUp).Methods("POST")
	router.HandleFunc("/loginForm", loginForm).Methods("GET")
	router.HandleFunc("/loginForm", login).Methods("POST")

	fmt.Println("Server Running Successfully")
	http.ListenAndServe("localhost:5000", router)
}

// function handling index.html
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Parsing template html file
	var tmpl, err = template.ParseFiles("index.html")
	//Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	var store = sessions.NewCookieStore([]byte("SESSION_ID"))
	session, _ := store.Get(r, "SESSION_ID")

	if session.Values["IsLogin"] != true {
		Data["IsLogin"] = false
	} else {
		Data["IsLogin"] = session.Values["IsLogin"].(bool)
		Data["Username"] = session.Values["Name"].(string)
	}

	rows, _ := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM public.tb_project;")

	var result []Project
	for rows.Next() {
		var each = Project{}

		var err = rows.Scan(&each.Id, &each.Title, &each.Start_date, &each.End_date, &each.Description, &each.Technologies, &each.Image)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		each.Format_start = each.Start_date.Format("12 May 2001")
		each.Format_end = each.End_date.Format("12 May 2001")

		//		if session.Values["IsLogin"] != true {
		//			each.IsLogin = false
		//		} else {
		//			each.IsLogin = session.Values ["IsLogin"].(bool)
		//		}

		result = append(result, each)
	}

	resp := map[string]interface{}{
		"Title":    Data,
		"Data":     Data,
		"Projects": result,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}

// function handling myproject.html
func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Parsing template html file
	var tmpl, err = template.ParseFiles("myproject.html")
	//Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

// function handling contactMe.html
func contactMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Parsing template html file
	var tmpl, err = template.ParseFiles("contactMe.html")
	//Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

// function handling myproiect-detail.html
func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//Parsing template html file
	var tmpl, err = template.ParseFiles("myproject-detail.html")
	//Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	ProjectDetail := Project{}

	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM public.tb_project WHERE id=$1", id).
		Scan(&ProjectDetail.Id, &ProjectDetail.Title, &ProjectDetail.Start_date, &ProjectDetail.End_date, &ProjectDetail.Description, &ProjectDetail.Technologies, &ProjectDetail.Image)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message: " + err.Error()))
		return
	}

	resp := map[string]interface{}{
		"Data":    Data,
		"Project": ProjectDetail,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	title := r.PostForm.Get("title")
	startDate := r.PostForm.Get("start-date")
	endDate := r.PostForm.Get("end-date")
	description := r.PostForm.Get("description")
	//java := r.PostForm.Get("Java")
	//php := r.PostForm.Get("Php")
	//laravel := r.PostForm.Get("Laravel")

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_project(name, start_date, end_date, description, technologies, image) VALUES($1, $2, $3, $4, '{NodeJs}', 'image.png')", title, startDate, endDate, description)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	//	var newProject = Project{
	//		Title:       title,
	//		Start_date:  startDate,
	//		End_date:    endDate,
	//		Description: description,
	//		NodeJs:      nodeJs,
	//		Java:        java,
	//		Php:         php,
	//		Laravel:     laravel,
	//	}

	//Projects = append(Projects, newProject)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func signupForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Parsing template html file
	var tmpl, err = template.ParseFiles("signUp.html")
	//Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 15)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO public.tb_user (name, email, password) VALUES($1, $2, $3)", name, email, passwordHash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	http.Redirect(w, r, "/loginForm", http.StatusMovedPermanently)
}

func loginForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Parsing template html file
	var tmpl, err = template.ParseFiles("login.html")
	//Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

func login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	user := User{}

	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	var store = sessions.NewCookieStore([]byte("SESSION_ID"))
	session, _ := store.Get(r, "SESSION_ID")

	session.Values["IsLogin"] = true
	session.Values["Name"] = user.Name
	session.Options.MaxAge = 10800

	session.AddFlash("Login Successful", "message")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
