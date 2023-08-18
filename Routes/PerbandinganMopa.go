package routes

import (
	handlers "mcdm/Handlers"
	mysql "mcdm/Pkg/Mysql"
	repositories "mcdm/Repositories"

	"github.com/gorilla/mux"
)

func PerbandinganMopaRoutes(r *mux.Router) {
	perbandinganMopaRepositori := repositories.RepositoryPerbandinganMopa(mysql.DB)
	h := handlers.HandlerperbandinganMopa(perbandinganMopaRepositori)

	r.HandleFunc("/perbandinganMopa", h.FindPerbandinganMopa).Methods("GET")
	r.HandleFunc("/perbandinganMopa/{kode}", h.GetPerbandinganMopaByKode).Methods("GET")
	r.HandleFunc("/perbandinganMopa", h.InsertPerbandinganMopa).Methods("POST")
	r.HandleFunc("/perbandinganMopa/edit/{kode}", h.UpdatePerbandinganMopa).Methods("PATCH")
	r.HandleFunc("/perbandinganMopa/delete/{kode}", h.DeletePerbandinganMopa).Methods("DELETE")
	// r.HandleFunc("/kriteria/lastCode", h.LastCode).Methods("POST")
}
