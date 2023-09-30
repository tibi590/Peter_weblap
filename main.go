package main

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
    Id string;
    Username string;
    Password string
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

    username := req.FormValue("input_name")
    password := req.FormValue("input_pass")

    var db_id string
    var db_user string
    var db_pass string

    fmt.Printf("Username: %v\n", username)
    fmt.Printf("Password: %v\n", password)

    err = db.QueryRow("SELECT id, username, password FROM users WHERE username=?", username).Scan(&db_id, &db_user, &db_pass)

    if err != nil {
        fmt.Println("ERROR: User Not Found")
        var Db_error = "Incorrect username"
        ftp.ExecuteTemplate(res, "login.html", Db_error)
        return
    }

    if password != db_pass {
        fmt.Println("Error: Incorrect Password")
        var Db_error = "Incorrect password"
        ftp.ExecuteTemplate(res, "login.html", Db_error)
        return
    }

    user := User{
        Id: db_id,
        Username: db_user,
        Password: db_pass,
    }

    fmt.Println("SUCCES: Redirecting To Home Page")
    ftp.ExecuteTemplate(res, "home.html", user)
}

func registerPage(res http.ResponseWriter, req *http.Request) {
    fmt.Println("####RegisterPage####")
    ftp.ExecuteTemplate(res, "register.html", nil)
}

func register(res http.ResponseWriter, req *http.Request) {
    fmt.Println("####Register####")
    
    if req.Method != "POST" {
        fmt.Println("ERROR: Not Post Request")
        http.ServeFile(res, req, ".pages/register.html")
    }

    username := req.FormValue("input_name")
    password := req.FormValue("input_pass")
    password_confirm := req.FormValue("input_pass_confirm")

    if password != password_confirm {
        fmt.Println("ERROR: Passwords Don't match")
        var Db_error = "Passwords don't match"
        ftp.ExecuteTemplate(res, "register.html", Db_error)
        return
    }

    var user string

    err = db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

    switch{
    case err == sql.ErrNoRows:
        _, err = db.Exec("INSERT INTO users(username, password) VALUES (?, ?)", username, password)
        if err != nil {
            http.Error(res, "SERVER ERROR: Unable To Create New User", 500)
            return
        }

        fmt.Println("SUCCES: New User Created")

        fmt.Println("SUCCES: Redirecting To Login Page")
        ftp.ExecuteTemplate(res, "login.html", nil)

    case err == nil:
        fmt.Println("ERROR: User Already Exists")
        var Db_error = "User already exists"
        ftp.ExecuteTemplate(res, "register.html", Db_error)
        return

    default:
        http.Error(res, "SERVER ERROR: Unable To Create New User", 500)
    }
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
    fmt.Println("SUCCES: Succesfully Pinged Database")

    http.HandleFunc("/", indexPage)
    http.HandleFunc("/login", loginPage)
    http.HandleFunc("/register", registerPage)
    http.HandleFunc("/db_login", login)
    http.HandleFunc("/db_register", register)

    http.Handle("/css/", http.StripPrefix("/css", css_files))
    http.Handle("/js/", http.StripPrefix("/js", js_files))

    http.ListenAndServe(":8080", nil)
}

