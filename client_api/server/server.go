package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type RouteRegister func() (router *mux.Router, port int)

type Logger interface {
	Infof(format string, args ...interface{})
	Error(args ...interface{})
}

func StartServer(l Logger, r RouteRegister) {
	router, port := r()
	http.Handle("/", router)
	l.Infof("server is started at %v\n", port)
	l.Error(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}


