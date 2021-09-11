package handlers

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type(
	UserController struct{
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponeWriter, r *http.Request, p httprouter.Parms){
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound) //404
	}

	oid := bson.ObjectHex(id)

	u := models.User{}

	if err := uc.Session.DB("mongo-golang").c("users").FindId(oid).one(&u); err != nil{  //c means the collections
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOk)
	fmt.Fprintf(w,"%s\n",uj)
}

func CreateUser(){}

func DeleteUser(){}