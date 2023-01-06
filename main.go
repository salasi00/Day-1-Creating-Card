package main

import (
	"fmt"
	"time"

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

type Project struct {
	Id int
	Title string
	StartDate string
	EndDate string
	Description string
	Technologies []string
	Duration string
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
	respData := map[string]interface{}{
		"Data" : Data,
		"Projects" : Projects,
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

	var projectDetail Project
	for index, data := range Projects{
		if index  == id {
			projectStartDate, _ := time.Parse("2006-01-02", data.StartDate)
			projectEndDate, _ := time.Parse("2006-01-02", data.EndDate)

			projectDetail = Project{
				Id: id,
				Title: data.Title,
				StartDate: projectStartDate.Format("02 Jan 2006"),
				EndDate: projectEndDate.Format("02 Jan 2006"),
				Description: data.Description,
				Technologies: data.Technologies,
				Duration: data.Duration,
			}
		}
	}

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
	

	

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _   := time.Parse("2006-01-02", dateEnd)

	timeDifference := endTime.Sub(startTime)

	year := int(timeDifference.Hours() / (12 * 30 * 24))
	month := int(timeDifference.Hours()/ (30 * 24))
	week := int(timeDifference.Hours() / (7 * 24))
	day := int(timeDifference.Hours() / 24)

	var duration string
	if year != 0 {
		duration = strconv.Itoa(year) + " Tahun"
	} else if month != 0 {
		duration = strconv.Itoa(month) + " Bulan"
	} else if week != 0 {
		duration = strconv.Itoa(week) + " Minggu"
	} else if day != 0 {
		duration = strconv.Itoa(day) + " Hari"
	}

	var newProject = Project{
		Title: title,
		StartDate: dateStart,
		EndDate: dateEnd,
		Description: description,
		Technologies: technologies,
		Duration: duration,
	}

	fmt.Println("Title: "  + r.PostForm.Get("title"))
	fmt.Println("Content: " + r.PostForm.Get("contents"))
	fmt.Println("Start Date: " + r.PostForm.Get("start-date"))
	fmt.Println("End Date: " + r.PostForm.Get("end-date"))
	fmt.Println("Technologies: ", r.Form["technologies"])

	Projects = append(Projects, newProject)

	fmt.Println(Projects)

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	Projects = append(Projects[:id], Projects[id+1:]...)
	
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

	var projectDetail Project

	for index, data := range Projects{
		if index == id {
			projectDetail = Project{
				Id: id,
				Title: data.Title,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Description: data.Description,
				Technologies: data.Technologies,
				Duration: data.Duration,
			}
		}
	}

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
	

	

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _   := time.Parse("2006-01-02", dateEnd)

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

	var newProject = Project{
		Title: title,
		StartDate: dateStart,
		EndDate: dateEnd,
		Description: description,
		Technologies: technologies,
		Duration: duration,
	}

	Projects[id] = newProject

	fmt.Println(Projects)

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