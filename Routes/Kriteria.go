package routes

import (
	handlers "mcdm/Handlers"
	mysql "mcdm/Pkg/Mysql"
	repositories "mcdm/Repositories"

	"github.com/gorilla/mux"
)

func KriteriaRoutes(r *mux.Router) {
	kriteriaRepositori := repositories.RepositoryKriteria(mysql.DB)
	h := handlers.HandlerKriteria(kriteriaRepositori)

	r.HandleFunc("/kriteria", h.FindKriteria).Methods("GET")
	r.HandleFunc("/kriteria/{kode}", h.GetKriteriaByKode).Methods("GET")
	r.HandleFunc("/kriteria", h.InsertKriteria).Methods("POST")
	r.HandleFunc("/kriteria/edit/{kode}", h.UpdateKriteria).Methods("PATCH")
	r.HandleFunc("/kriteria/delete/{kode}", h.DeleteKriteria).Methods("DELETE")
	r.HandleFunc("/kriteria/lastCode", h.LastCode).Methods("POST")
}
