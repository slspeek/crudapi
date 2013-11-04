package main

import (
	"github.com/gorilla/mux"
	"github.com/slspeek/crudapi"
	"github.com/slspeek/crudapi/storage/mongo"
	"github.com/slspeek/go_httpauth"
	"labix.org/v2/mgo"
	"log"
	"net/http"
)

type MyGuard struct {
	BasicAuth *go_httpauth.BasicServer
}

func NewMyGuard() MyGuard {
	return MyGuard{go_httpauth.NewBasic("MyRealm", func(user string, realm string) string {
		// Replace this with a real lookup function here
		return "ape"
	})}
}

func (g MyGuard) Authenticate(resp http.ResponseWriter, req *http.Request) (bool, string, string) {
	auth, username := g.BasicAuth.Auth(resp, req)
	log.Println("allowed:", auth, "username:", username)
	return auth, username, ""
}

func (g MyGuard) Authorize(client string, action crudapi.Action, urlVars map[string]string) (ok bool, errorMessage string) {
	log.Println("urlVars:", urlVars, "client:", client, "Action:", action)
	return true, ""
}

func hello(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello there!"))
}

func storage() *mongo.MongoStorage {
	s, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return mongo.NewMongoStorage(s, "test")
}

func main() {
	// storage

	s := storage()

	// router
	r := mux.NewRouter()

	// mounting the API
	crudapi.MountAPI(r.Host("localhost").Subrouter(), s, NewMyGuard())

	// custom handler
	r.HandleFunc("/", hello)

	// start listening
	log.Println("server listening on localhost:8080")
	log.Println("API on api.localhost:8080/v1/")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}
}
