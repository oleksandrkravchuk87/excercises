package video

import (
	"net/http"
)

// FormToUser -- fills a User struct with submitted form data
// params:
// r - request reader to fetch form data or url params (unused here)
// returns:
// User struct if successful
// array of strings of errors if any occur during processing
func FormToVideo(r *http.Request) (*Video, []string) {
	var video Video
	var errStr string
	var errs []string

	video.Title, errStr = processFormField(r, "title")
	errs = appendError(errs, errStr)
	video.Tag, errStr = processFormField(r, "tag")
	errs = appendError(errs, errStr)

	return &video, errs
}

func appendError(errs []string, errStr string) []string {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}

func processFormField(r *http.Request, field string) (string, string) {
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing '" + field + "' parameter, cannot continue"
	}
	return fieldData, ""
}
