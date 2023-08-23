package routes

import (
	handlers "mcdm/Handlers"
	mysql "mcdm/Pkg/Mysql"
	repositories "mcdm/Repositories"

	"github.com/gorilla/mux"
)

func PerbandinganCriteriaAhpRoutes(r *mux.Router) {
	perbandinganCriteriaAhpRepositori := repositories.RepositoryPerbandinganCriteriaAhp(mysql.DB)
	h := handlers.Handlerperbandingancriteriaahp(perbandinganCriteriaAhpRepositori)

	r.HandleFunc("/perbandinganCriteriaAhp", h.FindPerbandinganCriteriaAhp).Methods("GET")
	r.HandleFunc("/perbandinganCriteriaAhp/{kode}", h.GetPerbandinganCriteriaAhpByKode).Methods("GET")
	r.HandleFunc("/perbandinganCriteriaAhp", h.InsertPerbandinganCriteriaAhp).Methods("POST")
	r.HandleFunc("/perbandinganCriteriaAhp/edit/{kode}", h.UpdatePerbandinganCriteriaAhp).Methods("PATCH")
	r.HandleFunc("/perbandinganCriteriaAhp/delete/{kode}", h.DeletePerbandinganCriteriaAhp).Methods("DELETE")
	// r.HandleFunc("/kriteria/lastCode", h.LastCode).Methods("POST")
}
