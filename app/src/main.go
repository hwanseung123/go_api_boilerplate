package main

import (
   "fmt"
   "net/http"
   "os"
   _"github.com/gin-gonic/gin"
   "database/sql"
   "log"
   "gin-app/internal/handlers"
   _"github.com/lib/pq"
   _"github.com/golang-migrate/migrate/v4"
   _"github.com/golang-migrate/migrate/v4/database/postgres"
   _"github.com/golang-migrate/migrate/v4/source/file"
)

type User struct {
    Id int `json: "id"`
    Name string `json: "name"`
}

var db *sql.DB

func initDB() {
  conStr := os.Getenv("DATABASE_URL")
  if conStr == "" {
    log.Fatalf("DATABASE is empty")
  }
  var err error
  // sql.Open => This is essential for the code : modernc.org/sqlite
  db, err = sql.Open("postgres",conStr)

  if err != nil || db.Ping() != nil {
    log.Println("----------------- nil error ----------------")
    panic(err.Error())
  }
}

func main(){
   initDB()
   mux := http.NewServeMux()
   // "/" => route
   mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
     fmt.Println(w,"Hello world")
   })
   // "/insert" => route
   mux.HandleFunc("/insert",handlers.Insert(db))
   log.Println("server 8080: starting")
   // "/delte" => route
   mux.HandleFunc("/delete",handlers.Delete(db))
   // "/get" => route
   mux.HandleFunc("/get",handlers.Get(db))
   // "/put" => route
   mux.HandleFunc("/put",handlers.Put(db))

   err := http.ListenAndServe(":8080",mux)
   if err != nil {
       log.Fatalf("failed server...",err)
   }
}
