package main

import (
   "log"
   "os"
   gomail "gopkg.in/mail.v2"
)


func Gmail() {
   // Crate a message
   message := gomail.NewMessage()
 
   myEmail := "hwnia285@gmail.com"
   password := os.Getenv("PASSWORD")
   // Set email headers

   if password == "" {
      log.Println(password)
      log.Fatal("empty password")
   }
   message.SetHeader("From",myEmail)
   message.SetHeader("To",myEmail)
   message.SetBody("text/html","<h1>gokg.in buger</h1>")

   dialer := gomail.NewDialer("smtp.gmail.com",587,myEmail,password)
   if err := dialer.DialAndSend(message); err != nil {
      log.Println("Error:",err)
   }else {
      log.Println("HTML,Email sent successfuly")
    }
}
