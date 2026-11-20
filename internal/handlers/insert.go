package handlers

import (
   "fmt"
   "database/sql"
   "net/http"
   "encoding/json"
   _"github.com/gin-gonic/gin"
)

type UserRequest struct {
   Id int `json:"id"`
   Name string `json:"name"`
}

func Insert(db *sql.DB) http.HandlerFunc {
     // DB Clouser pattern
    insertUser := func(name string) error{
      fmt.Println(name)
      query := `INSERT INTO users (name) VALUES ($1)`
     // panin error
      _, err := db.Exec(query,name)

      if err != nil {
         panic(err)
      }
      return err
    }
     // HTTP POST
   return func (w http.ResponseWriter , r *http.Request) {
     // handlering panic error
     if db == nil {
        http.Error(w, "db is nil", http.StatusInternalServerError)
        return
     }
     if r.Method != http.MethodPost {
        http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
        return
    }
    // requst data JSON Parsing
    var user UserRequest
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil || user.Name == "" {
      fmt.Println("JSON Parsing error",err)
      http.Error(w,"Invaild request Body" , http.StatusBadRequest)
      return
    }

    // db insert
    // panic error
    err = insertUser(user.Name)
    if err != nil {
       http.Error(w,"Failed to insert user" , http.StatusInternalServerError)
       return
    }

    // http status code ok 200  => request header
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User insert successfuly"))
  }
}
