package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint8
	Money                 int16
	Avg_grades, Happiness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is:%s.He is %d and he have money"+
		" equal: %d", u.Name, u.Age, u.Money)
}
func (u *User) NewName(newName string) {
	u.Name = newName
}

func home_page(page http.ResponseWriter, r *http.Request) {
	Bob := User{Name: "Bob", Age: 21, Money: 38, Avg_grades: 3.5, Happiness: 0.8}
	//Bob.NewName("Alex")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(page, Bob)
}
func contacts_page(w http.ResponseWriter, r *http.Request) {}
func handleFunc() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}
func main() {
	handleFunc()
}
