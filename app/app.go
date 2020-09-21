package app

import (
	"net/http"	
	"github.com/sanju-png/web/userc.go"
)

func startapp(){
http.HandleFunc("/users",controllers.GetUser)

if err!=http.ListenAndServe(":8080", handler: nil); err != nil{
	panic(err)
}
}