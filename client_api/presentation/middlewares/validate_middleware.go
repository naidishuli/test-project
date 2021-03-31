package middlewares

import (
	"bytes"
	"client_api/presentation/response"
	"client_api/presentation/routes/schemas"
	"encoding/json"
	"github.com/gorilla/context"
	"io/ioutil"
	"net/http"
)

func ValidateBody(next http.HandlerFunc, validate schemas.Schema) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		bodyBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			response.ErrorResponse(res, req, err)
			return
		}
		req.Body.Close()
		context.Set(req, "bodyBytes", bodyBytes)

		encoder := json.NewDecoder(bytes.NewBuffer(bodyBytes))

		data, err := validate(encoder)
		if err != nil {
			response.ErrorResponse(res, req, err)
			return
		}

		context.Set(req, "payload", data)
		next(res, req)
	}
}
