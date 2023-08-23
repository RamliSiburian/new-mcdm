package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	KriteriaRoutes(r)
	AlternatifRoutes(r)
	AuthRoutes((r))
	PerbandinganAhpRoutes(r)
	PerbandinganCriteriaAhpRoutes(r)
	PerbandinganMopaRoutes(r)
}
