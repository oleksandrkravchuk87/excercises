package main

import (
	"encoding/json"
	"log"
	"net/http"

	"excercises/cassandratest/cassandra"

	"excercises/cassandratest/video"

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
	router.HandleFunc("/videos/new", video.Post).Methods(http.MethodPost)
	router.HandleFunc("/videos", video.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/videos", video.Delete).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8182", router))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
