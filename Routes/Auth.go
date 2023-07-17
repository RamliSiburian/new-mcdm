package routes

import (
	handlers "mcdm/Handlers"
	mysql "mcdm/Pkg/Mysql"
	midleware "mcdm/Pkg/Mysql/Midleware"
	repositories "mcdm/Repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/check-auth", midleware.Auth(h.CheckAuth)).Methods("GET")
}
