package app
import (
	"net/http"
	u "lens/utils"
	"strings"
	"go-contacts/models"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"context"
	"fmt"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(w http.ResponseWriter, r *http.Request){
		notAuth := []string{"/api/user/new","/api/user/login"}
		requestPath := r.URL.Path //current request path

		for _, value := range notAuth{
			if value == requestPath{
				next.ServeHTTP(w,r)
				return
			}
		}
	}
	response := make(map[string] interface{})
	tokenHeader := r.Header.Get("Authorization")//get token from header
	
	if tokenHeader == "" {
		response = u.Message("01", false, "Missing auth token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}

	split := strings.Split(tokenHeader," ")//split Bearer and Token
	if len(split) != 2{
		response = u.Message("01", false, "Invalid format auth Token")
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}

	tokenPart := split[1] //token
	tk := &models.Token{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token)(interface{}, error){
		return []byte(os.Getenv("token_password")), nil
	})
	if err != nil{
		response = u.Message("01", false, "Malformed authentication token")
		w.Header().Add("Content-type","application/json")
		u.Respond(w, response)
		return
	}

	if !token.Valid{ //Token is invalid, maybe not signed on this server
		response = u.Message("01", false, "Invalid Token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}

//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
	fmt.Sprintf("User %",tk.Username) //for monitoring
	ctx := context.WithValue(r.Context(), "user", tk.UserId)
	r = r.WithContext(ctx)
	next.ServeHTTP(w, r)
	});
}
