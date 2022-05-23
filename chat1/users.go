package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
	}
	temp.Execute(w, "Sign Up")
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/form.html")
	if err != nil {
		fmt.Println(err)
	}
	temp.Execute(w, "Sign In")
}
func Home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Println(err)
	}
	temp.Execute(w, "")
}

func getinfo(r *http.Request) (use *Users) {
	use = &Users{
		Email:    r.PostFormValue("Email"),
		Username: r.PostFormValue("Username"),
		Password: r.PostFormValue("Password"),
	}
	return
}

func SignI(w http.ResponseWriter, r *http.Request) {
	use := getinfo(r)
	if ((DefaultService.FindUser(use) != nil)){
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	} else {
		session, _ := store.Get(r, "session")
		session.Values["Email"] = use.Email
		session.Save(r, w)
		http.Redirect(w, r, "/messages", http.StatusSeeOther)
	}
}
func SignO(w http.ResponseWriter, r *http.Request) {
	use := getinfo(r)
	err := DefaultService.CreateUser(use)
	if (err != nil){
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		session, _ := store.Get(r, "session")
		session.Values["Email"] = use.Email
		fmt.Println(use.Username)
		session.Values["Username"] = use.Username
		session.Save(r, w)
		http.Redirect(w, r, "/messages", http.StatusSeeOther)
	}

}