package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/tawesoft/golib/v2/dialog"
)

// var Check Page

var tpl *template.Template

// session storing
var Store = sessions.NewCookieStore([]byte("session"))

// setting max age of the session
Store.Options = &sessions.Options{
    MaxAge: 10, // 10 seconds
}

// template init
func init() {
	tpl = template.Must(template.ParseGlob("static/*.html"))
}

type Page struct {
	Status bool
	//user input store the null interface
	Header1    interface{}
	IsLoggedin bool
	Valid      bool
}

var P = Page{
	Status: false,
}

// user data
var userData = map[string]string{
	"username": "abhinand",
	"password": "123",
}

// login function
func login(w http.ResponseWriter, r *http.Request) {
	//clearing
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	//checking the middle ware if user logined or not
	ok := Middleware(w, r)

	if ok {
		//redirecting
		http.Redirect(w, r, "/login-submit", http.StatusSeeOther)
		return
	}
	//if not logined
	P.Valid = Middleware(w, r)
	filename := "login.html"
	err := tpl.ExecuteTemplate(w, filename, P) //redirecting to login page
	if err != nil {
		fmt.Println("error while parsing file", err)
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//clearing
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	if r.Method == "GET" {
		//clearing
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		//checking middleware
		ok := Middleware(w, r)
		if ok {
			http.Redirect(w, r, "/", http.StatusSeeOther) //root file redirect
			return
		}
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "there is an error parsing %v", err)
		return
	}
	//store the username and pass
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	//print the username and password

	fmt.Println("username : ", username)
	fmt.Println("password : ", password)
		if userData["username"] == username && userData["password"] == password && r.Method == "POST" {

		//session storing
		session, _ := Store.Get(r, "started")
		// setting max age of the session on per session basis
		session.Options.MaxAge = 10 // 10 seconds
		//storing the value
		session.Values["username"] = username

		//value moving to P.Hedder1
		P.Header1 = session.Values["username"]

		//print the value
		fmt.Println("Hedder 1 value : ", P.Header1)

		//session saving
		session.Save(r, w)

		//print the session value
		fmt.Println("session value : ", session)

		//clearing
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		//wrong allert
		dialog.Alert("wrong passwod")
		//if its wrong --
		http.Redirect(w, r, "/login", http.StatusSeeOther)

		return

	}

}

// logout function
func Logouthandler(w http.ResponseWriter, r *http.Request) {
	//clearing
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	session, _ := Store.Get(r, "started")
	session.Options.MaxAge = -1 // delete session
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
