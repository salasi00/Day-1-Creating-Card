package main

import (
	"context"
	"fmt"
	"time"

	"log"
	"net/http"
	"personal-web/connection"

	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title"	  : "personal web",
	"isLogin" : true,
}

type Project struct {
	Id 					int
	Title 				string
	StartDate 			time.Time
	EndDate 			time.Time
	Description 		string
	Technologies 		[]string
	Duration 			string
}

var Projects = []Project{
	{
		// Title: "Dumbways Mobile Apps - 2021",
		// StartDate: "2022-09-30",
		// EndDate: "2022-12-30",
		// Description: "App that used for dumbways student, it was deployed and can be downloaded on playstore. Happy download",
		// Technologies: []string{"nodeJs","reactJs","nextJs","typeScript"},
		// Duration: "3 Month",
	},
}

func main() {
	route := mux.NewRouter()

	connection.DatabaseConnection()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	route.HandleFunc("/home", home).Methods("GET").Name("home")
	route.HandleFunc("/contact-form", contactForm).Methods("GET")
	route.HandleFunc("/blog", blogForm).Methods("GET")
	route.HandleFunc("/blog", addBlog).Methods("POST")
	route.HandleFunc("/blog-details/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/delete-blogs/{id}", deleteBlog).Methods("GET")
	route.HandleFunc("/edit-blogs/{id}", editBlog).Methods("GET")
	route.HandleFunc("/edit-blogs/{id}", updateBlog).Methods("POST")

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
	var result []Project
	rows, _ := connection.Conn.Query(context.Background(), "SELECT id, title, content, start_date, end_date, technologies FROM tb_blog")	
	for rows.Next(){
		var each = Project{}
		var err =  rows.Scan(&each.Id, &each.Title, &each.Description, &each.StartDate, &each.EndDate, &each.Technologies)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		each.Duration = getTimeDifference(each.StartDate, each.EndDate)
		result = append(result, each)
	}

	
	respData := map[string]interface{}{
		"Data" : Data,
		"Projects" : result,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
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

	projectDetail := Project{}

	err = connection.Conn.QueryRow(context.Background(), "SELECT id, title, content, start_date, end_date, technologies FROM tb_blog WHERE id=$1", id).Scan(
			&projectDetail.Id, &projectDetail.Title, &projectDetail.Description, &projectDetail.StartDate, &projectDetail.EndDate, &projectDetail.Technologies,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message: " + err.Error()))
		return
	}
	projectDetail.Duration = getTimeDifference(projectDetail.StartDate, projectDetail.EndDate)

	respData := map[string]interface{}{
		"Data" : Data,
		"projectDetail" : projectDetail,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)

}

func addBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	title := r.PostForm.Get("title")
	dateStart := r.PostForm.Get("start-date")
	dateEnd := r.PostForm.Get("end-date")
	description := r.PostForm.Get("contents")
	technologies := r.Form["technologies"]
	
	_ , err = connection.Conn.Exec(context.Background(),
	"INSERT INTO tb_blog(title, content, start_date, end_date, technologies) VALUES ($1, $2, $3, $4, $5)", 
	title, description, dateStart, dateEnd, technologies)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("message: " + err.Error()))
			return
		}

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blog WHERE id=$1", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message: " + err.Error()))
		return
	}
	
	http.Redirect(w, r, "/home", http.StatusFound)

}

func editBlog(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/edit-blog.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	projectDetail := Project{}
	
	err = connection.Conn.QueryRow(context.Background(), "SELECT id, title, content, start_date, end_date, technologies FROM tb_blog WHERE id=$1", id).Scan(
			&projectDetail.Id, &projectDetail.Title, &projectDetail.Description, &projectDetail.StartDate, &projectDetail.EndDate, &projectDetail.Technologies,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message: " + err.Error()))
		return
	}
	projectDetail.Duration = getTimeDifference(projectDetail.StartDate, projectDetail.EndDate)

	respData := map[string]interface{} {
		"Data":  Data,
        "projectDetail": projectDetail,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func updateBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	title := r.PostForm.Get("title")
	dateStart := r.PostForm.Get("start-date")
	dateEnd := r.PostForm.Get("end-date")
	description := r.PostForm.Get("contents")
	technologies := r.Form["technologies"]


	_ , err = connection.Conn.Exec(context.Background(),
	"UPDATE public.tb_blog SET title=$1, content=$2, start_date=$3, end_date=$4, technologies=$5 WHERE id=$6", 
	title, description, dateStart, dateEnd, technologies, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("message: " + err.Error()))
			return
		}

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

func getTimeDifference(startTime, endTime time.Time) string {

	timeDifference := endTime.Sub(startTime)

	year := int(timeDifference.Hours() / (12 * 30 * 24))
	month := int(timeDifference.Hours()/ (30 * 24))
	week := int(timeDifference.Hours() / (7 * 24))
	day := int(timeDifference.Hours() / 24)

	var duration string
	if year != 0 {
		duration = "durasi - " + strconv.Itoa(year) + " Tahun"
	} else if month != 0 {
		duration = "durasi - " +strconv.Itoa(month) + " Bulan"
	} else if week != 0 {
		duration = "durasi - " + strconv.Itoa(week) + " Minggu"
	} else if day != 0 {
		duration = "durasi - " + strconv.Itoa(day) + " Hari"
	} else {
		duration = "durasi - 0 Hari"
	}
	return duration
}
