package router

import (
	"github.com/MahmudovMZ/url-inspector/internal/transport/handler"
	"github.com/MahmudovMZ/url-inspector/internal/transport/middleware"
	"github.com/gorilla/mux"
)

func Rout() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.ValidateURL)
	r.HandleFunc("/inspect", handler.Inspect).Methods("GET")
	// inspect := r.PathPrefix("/inspect")
	return r
}
