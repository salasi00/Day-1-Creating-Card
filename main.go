package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title"	  : "personal web",
	"isLogin" : true,
}

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	route.HandleFunc("/home", home).Methods("GET").Name("home")
	route.HandleFunc("/contact-form", contactForm).Methods("GET")
	route.HandleFunc("/blog", blogForm).Methods("GET")
	route.HandleFunc("/blog", addBlog).Methods("POST")
	route.HandleFunc("/blog-details/{id}", blogDetail).Methods("GET")

	fmt.Println("Server is running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

func blogForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/blog.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var tmpl, err = template.ParseFiles("views/blog-details.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message: " + err.Error()))
		return
	}
	resp := map[string]interface{}{
		"Data" : Data,
		"Id" : id,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)

}

func addBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title: "  + r.PostForm.Get("title"))
	fmt.Println("Content: " + r.PostForm.Get("contents"))
	fmt.Println("Start Date: " + r.PostForm.Get("start-date"))
	fmt.Println("End Date: " + r.PostForm.Get("end-date"))
	fmt.Println("Technologies: ", r.Form["technologies"])
	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func contactForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/contact-form.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message: " + err.Error()))
		return
	}
	
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}
