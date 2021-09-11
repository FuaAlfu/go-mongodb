package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	handlers "go-mongodb/pkg/handlers"
)

func routes() {
	r := httprouter.New()
	//new session
	uc := handlers.NewUserController(getSession())
	r.GET("/user/:id",uc.GetUser)
	r.POST("/user",uc.CreateUser)
	r.DELETE("/user/:id",uc.DeleteUser)
	http.ListenAndServe("localhost:9090",r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27187")
	if err != nil{
		panic(err)
	}
	return s
}

func main() {
	routes()
}
