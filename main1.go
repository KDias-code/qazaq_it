package main

import(
	"fmt"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil{
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
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

func handleFunc(){
	http.HandleFunc("/", index)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/save_article/", save_article) 
	http.ListenAndServe(":8080", nil)
}

func main(){
	handleFunc()
}
