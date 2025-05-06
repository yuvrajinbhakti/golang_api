package handlers 
import {
	"github.com/go-chi/chi" 
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/yuvrajinbhakti/golang_api/internal/middleware"
}

//middleware is a function that gets called before the primary function which handles the endpoint
func Handler(r *chi.Mux){
//by r.use we can add middleware to the route 
//global middleware (it mean that this middleware will be applied to all the routes/endpoints)
r.Use(chimiddle.StripSlashes)

r.Route("/account",func(router chi.Router){
	router.Use(middleware.Authorization) // this middleware will check if use is authorized to access this data
 	router.Get("/coins",GetCoinsBalance)
})
 
}