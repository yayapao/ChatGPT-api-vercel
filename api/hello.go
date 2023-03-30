package handler

import (
	"fmt"
	"net/http"
	"time"
)

// ResponseCurrentTime returns the current time, you can access it by /api/hello
func ResponseCurrentTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	_, err := fmt.Fprintln(w, currentTime)
	if err != nil {
		return
	}
}
