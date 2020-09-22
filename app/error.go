package app

import (
	u "api-go-deploy-heroku/utils"
	"net/http"
)

var NotFoundHandler = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message("01", false, "This resources was not found on our server"))
		next.ServeHTTP(w, r)
	})
}
