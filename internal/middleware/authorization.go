package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/yuvrajinbhakti/golang_api/api"
	"github.com/yuvrajinbhakti/golang_api/internal/tools"
)

var UnAuthorizedError = errors.New("Invalid username or token") //custom unauthorized error 

func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error 
		if username == "" || token == ""{
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w,UnAuthorizedError)
			return
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}
		var loginDetails *tools.LoginDetails 
		loginDetails = (*database).GetUserLoginDetails(username)
		if (loginDetails == nil || (token != (*loginDetails).AuthToken)){
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w,UnAuthorizedError)
			return
		}
		next.ServeHTTP(w,r) //this function in the line or in handler function for the endpoint if there's no middlerware left
	})
}