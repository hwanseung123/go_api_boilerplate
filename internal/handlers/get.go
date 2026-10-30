package handlers

import (
    "fmt"
    "net/http"
    "database/sql"
    _"encoding/json"
)


type PersonGetResponse struct {
   Id int `json:"id"`
   Name string `json:"name"`
}


func Get(db *sql.DB) http.HandlerFunc{
    getUser := func() ([]PersonGetResponse,error) {
       query :=  `SELECT id,name FROM users`
       resp, err := db.Query(query) // get => db.Query
       defer resp.Close() // end db memeory close

       var users []PersonGetResponse
       for resp.Next() {
         var u PersonGetResponse
         if err = resp.Scan(&u.Id , &u.Name); err != nil {
            return nil, err
         }
          // users append (var users []Person) = (var users []Person,new user) 
          users = append(users,u)
       }
       // not error => nil
       /* res
           why it is nil?
           0 value => not int,float,string

         */
       return users, nil
    }
    return func(w http.ResponseWriter , r *http.Request){
       if r.Method != http.MethodGet {
         http.Error(w, "Method not allowed" , http.StatusMethodNotAllowed)
         return
       }
      users, err := getUser()
      if err != nil {
        http.Error(w, "Failed to get users", http.StatusInternalServerError)
        return
      }
      for index,user := range users {
          fmt.Printf("Index: %d , Id: %d , Name: %s\n",index,user.Id,user.Name)
      }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Users get successfuly"))
    }
}
