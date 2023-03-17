package main

import ("fmt"; "net/http"; "html/template")

type User struct{
	Name string
	Age uint16
	Money int16
	Avg_grades, Happiness float64
	Hobbies []string
}

func(u User) getAllInfo() string{
	return fmt.Sprintf("User name is: %s. He is %d years old and have %d $", u.Name, u.Age, u.Money)
}

func home_page(w http.ResponseWriter, r *http.Request){
	user1 := User{"Almaz", 22, -40, 3.3, 0.9, []string{"Saken", "Rauan", "Arman"}}
	// fmt.Fprintf(w, user1.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, user1)
}

func contact_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "87789524881")
}

func main ()  {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contact/", contact_page)
	http.ListenAndServe(":8080", nil)

}