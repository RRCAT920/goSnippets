package main

import (
	"goSnippets/filelistingserver/controller"
	"log"
	"net/http"
	"os"
)

type errorHandler func(http.ResponseWriter, *http.Request) error

func errorWrap(handler errorHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Error occurred handling request: %s\n", err)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			if handleUserErrorSucceed(writer, err) {
				return
			}
			handleHTTPError(writer, err)
		}
	}
}

func handleHTTPError(writer http.ResponseWriter, err error) {
	log.Printf("Error occurred handling request: %s\n", err)
	code := http.StatusOK
	switch {
	case os.IsNotExist(err):
		code = http.StatusNotFound
	case os.IsPermission(err):
		code = http.StatusForbidden
	default:
		code = http.StatusInternalServerError
	}
	http.Error(writer, http.StatusText(code), code)
}

func handleUserErrorSucceed(writer http.ResponseWriter, err error) bool {
	if userErr, ok := err.(userError); ok {
		log.Printf("Error occurred handling request: %s\n", userErr.Message())
		http.Error(writer,
			userErr.Message(),
			http.StatusBadRequest)
		return true
	}
	return false
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errorWrap(controller.ListFile))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
