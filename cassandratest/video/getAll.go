package video

import (
	"encoding/json"
	"excercises/cassandratest/cassandra"
	"net/http"

	"fmt"
	"time"

	"github.com/gocql/gocql"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	var videoList []Video
	m := map[string]interface{}{}
	fmt.Println("-------")
	query := "SELECT * FROM video_by_tag"
	iterable := cassandra.Session.Query(query).Iter()
	for iterable.MapScan(m) {
		videoList = append(videoList, Video{
			ID:        m["video_id"].(gocql.UUID),
			AddedDate: m["added_date"].(time.Time),
			Title:     m["title"].(string),
			Tag:       m["tag"].(string),
		})
		m = map[string]interface{}{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AllVideosResponse{Videos: videoList})

}
