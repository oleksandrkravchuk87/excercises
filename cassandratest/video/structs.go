package video

import (
	"time"

	"github.com/gocql/gocql"
)

// Video struct to hold profile data for our user
type Video struct {
	ID        gocql.UUID `json:"videoID"`
	AddedDate time.Time  `json:"addedDate"`
	Title     string     `json:"title"`
	Tag       string     `json:"tag"`
}

// GetUserResponse to form payload returning a single User struct
type GetVideoResponse struct {
	Video Video `json:"video"`
}

// AllUsersResponse to form payload of an array of User structs
type AllVideosResponse struct {
	Videos []Video `json:"videos"`
}

// NewUserResponse builds a payload of new user resource ID
type NewVideoResponse struct {
	ID gocql.UUID `json:"videoid"`
}

// ErrorResponse returns an array of error strings if appropriate
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
