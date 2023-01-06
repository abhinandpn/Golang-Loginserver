# SIgnuppage
Package: main

Imports:
- "fmt"
- "html/template"
- "net/http"
- "github.com/gorilla/sessions"
- "github.com/tawesoft/golib/v2 v2.6.1 // indirect"

Type: page

- Status: bool
- Header1: interface{}
- IsLoggedin: bool
- Valid: bool

Var:
- tpl: *template.Template
- Store: sessions.NewCookieStore([]byte("session"))
- P: page{
    Status: false
  }
- userData: map[string]string{
    "username":    "abhinand",
    "password": "123",
  }

Func: init
- Initializes the template by parsing the templates located in the "template/*.html" glob pattern.

Func: index
- Handles the "/" route.
- Executes the "index.html" template and sends the output to the http.ResponseWriter.
- Calls the Middleware function to check if the user is logged in.
- If the user is logged in, sets the page.Status value to true.

Func: login
- Handles the "/login" route.
- Executes the "login.html" template and sends the output to the http.ResponseWriter.
- Calls the Middleware function to check if the user is logged in.
- If the user is logged in, redirects to the "/login-submit" route.

Func: loginHandler
- Handles the "/login-submit" route.
- If the request method is GET, calls the Middleware function to check if the user is logged in.
- If the user is logged in, redirects to the "/" route.
- Parses the form data from the request.
- Gets the email and password values from the request form data.
- If the email and password match the ones in the userData map and the request method is POST, creates a new session and sets the "id" value to "AKSHAY".
- Sets the page.Header1 value to the session "id" value.
- Saves the session and redirects to the "/" route.
- If the email and password don't match or the request method is not POST, redirects to the "/login" route.

Func: logoutHandler
- Handles the "/logout" route.
- If the user is logged in (page.Status is true), gets the session and sets the MaxAge value to -1 to expire the session.
- Saves the session and redirects to the "/" route.
- If the user is not logged in (page.Status is false), redirects to the "/login" route.

Func: Middleware
- Takes an http.ResponseWriter and an http.Request as arguments.
- Gets the session and checks if the "id" value is set.
- If the "id" value is set, returns true.
- If the "id" value is not set, returns false.
