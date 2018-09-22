package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

var staticFilterHandler = func(h http.Handler, exts []string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			for _, ext := range exts {
				if filepath.Ext(r.URL.Path) == ext {
					h.ServeHTTP(w, r)
				}
			}
		})
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		log.Printf("Handler error: status code: %d, message: %s, underlying err: %#v",
			e.Code, e.Message, e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

func appErrorf(err error, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    http.StatusInternalServerError,
	}
}

func responseWithData(w http.ResponseWriter, status int, data interface{}) {
	var b bytes.Buffer
	if err := encode(&b, data); err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(b.Bytes())
}

func encode(w io.Writer, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}

func decode(r io.ReadCloser, data interface{}) error {
	return json.NewDecoder(r).Decode(&data)
}
