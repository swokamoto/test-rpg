package main

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"

    "github.com/uptrace/bun"
    "github.com/uptrace/bun/dialect/pgdialect"
    "github.com/uptrace/bun/driver/pgdriver"
)

var db *bun.DB

func main() {
    dsn := "postgres://user:pass@localhost:5432/mydb?sslmode=disable"
    sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
    db = bun.NewDB(sqldb, pgdialect.New())

    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/hello", helloHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("templates/index.templ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    data := struct {
        Title   string
        Heading string
    }{
        Title:   "testRPG",
        Heading: "Hello, World!",
    }
    if err := t.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}