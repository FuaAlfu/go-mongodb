package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func routes() {
	r := httprouter.New()
	//new session
	uc := handlers.NewUserController(getSession())
	r.GET("/")
	r.POST("")
	r.DELETE("")
}

func getSession() *mgo.Session {

}

func main() {
	routes()
}
