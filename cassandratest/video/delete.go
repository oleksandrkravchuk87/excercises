package video

import (
	"excercises/cassandratest/cassandra"
	"net/http"

	"fmt"

	"time"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	//var errs []string
	id := r.URL.Query().Get("id")
	tag := r.URL.Query().Get("tag")
	date := r.URL.Query().Get("date")

	tsAfter, err := time.Parse(time.RFC3339, date)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(date)
	fmt.Println(tsAfter)
	fmt.Println(id)
	fmt.Println(tag)
	err = cassandra.Session.Query(`DELETE FROM video_by_tag WHERE tag = ? AND added_date = ? AND video_id = ?`, tag, tsAfter, id).Exec()
	if err != nil {
		fmt.Println("errors", err)

	} else {

		fmt.Println("del")
	}
}
