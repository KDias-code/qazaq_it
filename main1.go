package main

import(
	"fmt"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Articles struct{
	Id uint16
	Anons, Full_text, Tittle string
}

var posts = []Articles{}
var showPost = Articles{}

func index(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil{
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3316)/golang")
		if err != nil{
			panic(err)
		}

		defer db.Close()
		
		res, err := db.Query("SELECT * FROM articles")
		if err != nil{
			panic(err)
		}

		posts = []Articles{}

		for res.Next(){
			var post Articles
			err = res.Scan(&post.Id, &post.Tittle, &post.Anons, &post.Full_text)
			if err != nil{
				panic(err)
			}
			posts = append(posts, post)
		}

	t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil{
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request){
	tittle := r.FormValue("tittle")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")
	// fmt.Fprintf(w, tittle, anons, full_text)

	if tittle == "" || anons == "" || full_text == ""{
		fmt.Fprintf(w, "One of the lines empty")
	} else{
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3316)/golang")
		if err != nil{
			panic(err)
		}

		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO articles (tittle, anons, full_text) VALUES('%s', '%s', '%s')", tittle, anons, full_text))
		if err != nil{
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func show_post( w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

	if err != nil{
		fmt.Fprintf(w, err.Error())
	}

	// w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3316)/golang")
		if err != nil{
			panic(err)
		}

		defer db.Close()

		res, err := db.Query(fmt.Sprintf("SELECT * FROM articles WHERE id = '%s'", vars["id"]))
		if err != nil{
			panic(err)
		}

		showPost = Articles{}

		for res.Next(){
			var post Articles
			err = res.Scan(&post.Id, &post.Tittle, &post.Anons, &post.Full_text)
			if err != nil{
				panic(err)
			}
			showPost = post
		}

	t.ExecuteTemplate(w, "show", showPost)
}

func contacts(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/contacts.html", "templates/header.html", "templates/footer.html")

	if err != nil{
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "contacts", nil)
}

func handleFunc(){
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")
	rtr.HandleFunc("/contacts", contacts).Methods("GET")

	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}

func main(){
	handleFunc()
}
