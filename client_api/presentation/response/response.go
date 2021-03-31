package response

import (
	. "client_api/app/errors"
	"client_api/domain/interfaces"
	"client_api/infrastructure/logger"
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var log interfaces.Logger

func init() {
	log = logger.NewLogger(logrus.Fields{})
}

func ErrorResponse(res http.ResponseWriter, r *http.Request, err error) {
	var status int
	res.Header().Set("content-type", "application/json")
	switch v := err.(type) {
	case *APIError:
		status = v.StatusCode()
		res.WriteHeader(v.StatusCode())
		_ = json.NewEncoder(res).Encode(v)
	default:
		status = http.StatusInternalServerError
		res.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(res).Encode(InternalServerError(err))
	}

	startTime := context.Get(r, "startExecTime").(time.Time)
	elapsed := time.Since(startTime)
	log.Warnf("[%s] %s %s %d\n", elapsed, r.Method, r.RequestURI, status)
}

func SuccessResponse(res http.ResponseWriter, r *http.Request, body interface{}) {
	res.Header().Set("content-type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(body)
	startTime := context.Get(r, "startExecTime").(time.Time)
	elapsed := time.Since(startTime)
	log.Infof("[%s] %s %s %d\n", elapsed, r.Method, r.RequestURI, 200)
}
