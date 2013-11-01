package main

import (
	"github.com/gorilla/mux"
	"github.com/sauerbraten/crudapi"
	"github.com/sauerbraten/crudapi/storage/mongo"
	"labix.org/v2/mgo"
	"log"
	"net/http"
)

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
	crudapi.MountAPI(r.Host("localhost").Subrouter(), s, nil)

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
