package main 

import {
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/yuvrajinbhakti/golang_api/internal/handlers"
	log "github.com/sirupsen/logrus"  // for logging
}
// go mod tidy -> to install all the dependencies

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
    fmt.Println("Starting GO API Service...")
	fmt.Println(`
GO API  
	`)
	err := http.ListenAndServe("localhost:8000",r)
	if err != nil{
		log.Error(err)
	}
}