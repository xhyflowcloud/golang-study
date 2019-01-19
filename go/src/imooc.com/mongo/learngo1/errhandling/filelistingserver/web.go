package main

import (
	"imooc.com/mongo/learngo1/errhandling/filelistingserver/handle"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func wrapperHandler(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handle.HandleFileList(writer, request)
		switch {
		case os.IsNotExist(err):
			http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func main() {
	http.HandleFunc("/list/", wrapperHandler(handle.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
