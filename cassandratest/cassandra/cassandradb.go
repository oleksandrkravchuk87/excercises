package cassandra

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "killrvideo"
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("cassandra initialised")
}
