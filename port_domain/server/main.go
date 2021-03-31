package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"port_domain/infrastructure/logger"
)

func main() {
	l := logger.NewLogger(logrus.Fields{})
	StartServer(l, buildRouter)
}

func buildRouter() {
}

func ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

