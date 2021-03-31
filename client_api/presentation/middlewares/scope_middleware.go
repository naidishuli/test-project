package middlewares

import (
	"client_api/app"
	"github.com/gorilla/context"
	"net/http"
)

func AddScope(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		scope := app.NewScope()

		//todo the below line for a full implementation
		//user, ok := context.Get(r, "user").(models.User)
		//if !ok {
		//	response.ErrorResponse(w, r, fmt.Errorf("server error"))
		//	return
		//}
		//scope.SetUser(user)

		context.Set(r, "appScope", scope)
		next.ServeHTTP(w, r)
	})
}
