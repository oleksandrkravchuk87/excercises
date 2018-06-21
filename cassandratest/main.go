package main

import (
	"encoding/json"
	"log"
	"net/http"

	"excercises/cassandratest/cassandra"

	"excercises/cassandratest/user"

	"github.com/gorilla/mux"
)

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func main() {

	CassandraSession := cassandra.Session
	defer CassandraSession.Close()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", heartbeat).Methods(http.MethodGet)
	router.HandleFunc("/users/new", user.Post).Methods(http.MethodPost)
	router.HandleFunc("/users", user.GetAll).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8181", router))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
