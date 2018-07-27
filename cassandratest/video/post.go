package video

import (
	"encoding/json"
	"excercises/cassandratest/cassandra"
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var errs []string
	var gocqlUuid gocql.UUID

	video, errs := FormToVideo(r)

	created := false
	if len(errs) == 0 {
		fmt.Println("creating a new user")

		gocqlUuid = gocql.TimeUUID()
		if err := cassandra.Session.Query(`
			INSERT INTO video_by_tag (video_id, added_date, title, tag) 
			VALUES ( ?, toUnixTimestamp(now()), ?, ?)`,
			gocqlUuid, video.Title, video.Tag).Exec(); err != nil {
			errs = append(errs, err.Error())
		} else {
			created = true
		}

	}
	if created {
		fmt.Println("user_id", gocqlUuid)
		json.NewEncoder(w).Encode(NewVideoResponse{ID: gocqlUuid})
	} else {
		fmt.Println("errors", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
