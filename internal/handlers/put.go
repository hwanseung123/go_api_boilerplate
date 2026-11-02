package handlers

import (
    "fmt"
    "log"
    "database/sql"
    "net/http"
    "encoding/json"
)


type UserPutRequest struct {
   Id  int  `json:"id"`
   Name  string `json:"name"`
}


func Put(db *sql.DB) http.HandlerFunc{
  putUser := func(id int ,name string) error{
      query := `UPDATE users SET name = ? WHERE id = ?`
      _, err := db.Exec(query,name,id)
      return err
  }

  return func (w http.ResponseWriter , r *http.Request){
    if r.Method != http.MethodPut {
       http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
       return
    }

    var user UserPutRequest

    err := json.NewDecoder(r.Body).Decode(&user)
    err = putUser(user.Id,user.Name)
    log.Println(user.Id)
    if err != nil {
       fmt.Println("decoding error",err)
    }

    if err != nil {
      http.Error(w,"Failed to update user" , http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User update successfuly"))

  }
}
