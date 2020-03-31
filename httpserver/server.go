package httpserver

import (
	"net/http"
	"regexp"
	"strconv"
)

// New creates a new server with the codes as paths
func New(addr string, codes []int) *http.Server {
	r := regexpHandler{}

	for _, i := range codes {
		r.HandleFunc(regexp.MustCompile("/"+strconv.Itoa(i)), newHandleFunc(i))
	}

	server := http.Server{
		Addr:    addr,
		Handler: &r,
	}

	return &server
}

func newHandleFunc(httpCode int) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(httpCode)
	}
}
