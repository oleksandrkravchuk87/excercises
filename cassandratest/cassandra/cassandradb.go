package cassandra

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("172.17.0.2", "172.17.0.3")
	cluster.Keyspace = "streamdemoapi"
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("cassandra initialised")
}
