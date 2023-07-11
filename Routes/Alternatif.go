package routes

import (
	handlers "mcdm/Handlers"
	mysql "mcdm/Pkg/Mysql"
	repositories "mcdm/Repositories"

	"github.com/gorilla/mux"
)

func AlternatifRoutes(r *mux.Router) {
	alternatifRepositori := repositories.RepositoryAlternatif(mysql.DB)
	h := handlers.HandlerAlternatif(alternatifRepositori)

	r.HandleFunc("/alternatif", h.FindAlternatif).Methods("GET")
	r.HandleFunc("/alternatif/lastCode", h.LastCode).Methods("POST")
	r.HandleFunc("/alternatif", h.InsertAlternatif).Methods("POST")
	r.HandleFunc("/alternatif/{kode}", h.GetAlternatifByKode).Methods("GET")
	r.HandleFunc("/alternatif/delete/{kode}", h.DeleteAlternatif).Methods("DELETE")
	r.HandleFunc("/alternatif/edit/{kode}", h.UpdateAlternatif).Methods("PATCH")

}
