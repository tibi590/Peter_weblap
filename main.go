package main

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
    id string;
    username string;
    password string
}

var ftp *template.Template
var db *sql.DB
var err error

func indexPage(res http.ResponseWriter, req *http.Request) {
    fmt.Println("####IndexPage####")
    ftp.ExecuteTemplate(res, "index.html", nil)
}

func loginPage(res http.ResponseWriter, req *http.Request) {
    fmt.Println("####LoginPage####")
    ftp.ExecuteTemplate(res, "login.html", nil)
}

func login(res http.ResponseWriter, req *http.Request) {
    fmt.Println("####Login####")

    if req.Method != "POST" {
        fmt.Println("ERROR: Not Post Request")
        http.ServeFile(res, req, ".pages/login.html")
    }

    username := req.FormValue("input_user")
    password := req.FormValue("input_pass")

    var db_id string
    var db_user string
    var db_pass string

    err := db.QueryRow("SELECT id, username, password FROM users WHERE username=?", username).Scan(&db_id, &db_user, &db_pass)

    if err != nil {
        fmt.Println("ERROR: Unable To Execute Query")
        http.Redirect(res, req, "/login", 301)
        return
    }

    if password != db_pass {
        fmt.Println("Error: Incorrect Password")
        http.Redirect(res, req, "/login", 301)
        return
    }

    user := User{db_id, db_user, db_pass}

    fmt.Println("SUCCES: Redirecting To Home Page")
    ftp.ExecuteTemplate(res, "home.html", user)
}

func registerPage(res http.ResponseWriter, req *http.Request) {
    fmt.Println("####RegisterPage####")
    ftp.ExecuteTemplate(res, "register.html", nil)
}

func main() {
    ftp, _ = ftp.ParseGlob("pages/*.html")
    css_files := http.FileServer(http.Dir("./css"))
    js_files := http.FileServer(http.Dir("./js"))

    db, err = sql.Open("mysql", "admin:admin@/peter")
    if err != nil {
        fmt.Println("ERROR: Unable To Connect To Database")
        panic(err.Error())
    }
    fmt.Println("SUCCES: Connected To Database")
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    http.HandleFunc("/", indexPage)
    http.HandleFunc("/login", loginPage)
    http.HandleFunc("/register", registerPage)
    http.HandleFunc("/db_login", login)

    http.Handle("/css/", http.StripPrefix("/css", css_files))
    http.Handle("/js/", http.StripPrefix("/js", js_files))

    http.ListenAndServe(":8080", nil)
}

