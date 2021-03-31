package main

import (
	"client_api/app/config"
	"client_api/infrastructure/logger"
	"client_api/presentation/middlewares"
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	l := logger.NewLogger(logrus.Fields{})
	StartServer(l, buildRouter)
}

func buildRouter() (*mux.Router, int) {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router), execTime)
	router.HandleFunc("/ping", ping).Methods("GET", "HEAD")

	routerV1 := router.PathPrefix("/v1").Subrouter()
	routerV1.Use(middlewares.AddScope)

	//todo for full implementation add Authenticate middleware
	securedRouterV1 := router.PathPrefix("/v1").Subrouter()
	//securedRouterV1.Use(middlewares.Authenticate)
	securedRouterV1.Use(middlewares.AddScope)
	return router, config.Env.Port
}

func ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func execTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		context.Set(r, "startExecTime", start)
		next.ServeHTTP(w, r)
	})
}
