package routes

import (
	handlers "mcdm/Handlers"
	mysql "mcdm/Pkg/Mysql"
	repositories "mcdm/Repositories"

	"github.com/gorilla/mux"
)

func PerbandinganAhpRoutes(r *mux.Router) {
	perbandinganAhpRepositori := repositories.RepositoryPerbandinganAhp(mysql.DB)
	h := handlers.Handlerperbandinganahp(perbandinganAhpRepositori)

	r.HandleFunc("/perbandinganAhp", h.FindPerbandinganAhp).Methods("GET")
	r.HandleFunc("/perbandinganAhp/{kode}", h.GetPerbandinganAhpByKode).Methods("GET")
	r.HandleFunc("/perbandinganAhp", h.InsertPerbandinganAhp).Methods("POST")
	r.HandleFunc("/perbandinganAhp/edit/{kode}", h.UpdatePerbandinganAhp).Methods("PATCH")
	r.HandleFunc("/perbandinganAhp/delete/{kode}", h.DeletePerbandinganAhp).Methods("DELETE")
	// r.HandleFunc("/kriteria/lastCode", h.LastCode).Methods("POST")
}
