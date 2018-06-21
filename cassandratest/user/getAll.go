package user

import (
	"encoding/json"
	"excercises/cassandratest/cassandra"
	"net/http"

	"fmt"

	"github.com/gocql/gocql"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	var userList []User
	m := map[string]interface{}{}
	fmt.Println("-------")
	query := "SELECT id, age, firstname, lastname, city, email FROM users"
	iterable := cassandra.Session.Query(query).Iter()
	for iterable.MapScan(m) {
		userList = append(userList, User{
			ID:        m["id"].(gocql.UUID),
			Age:       m["age"].(int),
			FirstName: m["firstname"].(string),
			LastName:  m["lastname"].(string),
			Email:     m["email"].(string),
			City:      m["city"].(string),
		})
		m = map[string]interface{}{}
	}
	json.NewEncoder(w).Encode(AllUsersResponse{Users: userList})
}
