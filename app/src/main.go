package main

import (
   "fmt"
   "net/http"
   _"github.com/gin-gonic/gin"
   "database/sql"
   "log"
   _"modernc.org/sqlite"
   "gin-app/internal/handlers"
)

type User struct {
    Id int `json: "id"`
    Name string `json: "name"`
}


var db *sql.DB
func initDB() {
  var err error
  // sql.Open => This is essential for the code : modernc.org/sqlite 
  db,err = sql.Open("sqlite","demo.db")
  if err != nil{
       log.Fatalf("Failed to open database: %v",err)
  }

  createTableSQL := `CREATE TABLE IF NOT EXISTS users (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
     name TEXT NOT NULL
  );`

  if _, err := db.Exec(createTableSQL); err != nil {
       log.Fatalf("Failed to create table: %v",err)
  }

  fmt.Println("success table@@@@")
}

/*
func helloHandler(w http.ResponseWriter, r *http.Request){
   if r.URL.Path != "/" {
      http.NotFound(w,r)
      return
  }
   fmt.Fprint(w,"Hello World")
}
*/


func main(){
   initDB()
   // http.HandleFunc("/",helloHandler)
   // err := http.ListenAndServe(":8080",nil)
   // fmt.Println("Server 8080")
   // router := gin.Default()
   mux := http.NewServeMux()
   // "/" => route
   mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
     fmt.Println(w,"Hello world")
   })
   // "/insert" => route
   mux.HandleFunc("/insert",handlers.Insert(db))
   log.Println("server 8080: starting")

   err := http.ListenAndServe(":8080",mux)
   if err != nil {
       log.Fatalf("failed server...",err)
   }
}
