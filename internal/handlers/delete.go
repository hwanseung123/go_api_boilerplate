package handlers

import (
  _"fmt"
  "log"
  "net/http"
  _"encoding/json"
  "database/sql"
)

func Delete(db *sql.DB) http.HandlerFunc {

    return func (w http.ResponseWriter , r *http.Request) {
       if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
      }
       name := r.URL.Query().Get("name")
       if name == "" {
           http.Error(w, "Missing 'name' query parameter",http.StatusBadRequest)
           return
       }

      log.Println("name:",name)
      _, err := db.Exec("DELETE FROM users WHERE name = ?",name)

      if err != nil {
         http.Error(w, "Database Error", http.StatusInternalServerError)
         return
      }
      w.WriteHeader(http.StatusOK)
      w.Write([]byte("Deleted name: "+name))
    }
}


