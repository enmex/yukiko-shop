package response

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSON marshals 'v' to JSON, automatically escaping HTML and setting the
// Content-Type as application/json.
func JSON(w http.ResponseWriter, status int, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Reply(w, status, buf.Bytes())
}

func Reply(w http.ResponseWriter, status int, content []byte) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	_, _ = w.Write(content)
}

func EmptyJSON(w http.ResponseWriter, status int) {
	buf := &bytes.Buffer{}
	Reply(w, status, buf.Bytes())
}
