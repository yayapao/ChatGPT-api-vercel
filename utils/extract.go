package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func ExtractFromBody(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

// Stream sends a streaming response and returns a boolean
// indicates "Is client disconnected in middle of stream"
func Stream(w http.ResponseWriter, step func(w io.Writer) bool) bool {
	clientGone := w.(http.CloseNotifier).CloseNotify()
	for {
		select {
		case <-clientGone:
			return true
		default:
			keepOpen := step(w)
			w.(http.Flusher).Flush()
			if !keepOpen {
				return false
			}
		}
	}
}
