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
)

type User struct {
    Id int `json: "id"`
    Name string `json: "name"`
}

var db *sql.DB

func initDB() {
  var err error
  conStr := os.Getenv("DATABASE_URL")
  // sql.Open => This is essential for the code : modernc.org/sqlite 
  db,err = sql.Open("postgres",conStr)
  if err != nil {
    panic(err)
  }

  if err != nil{
       log.Fatalf("Failed to open database: %v",err)
  }
  fmt.Println("success table@@@@")
}

func main(){
   initDB()
   Gmail()
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
