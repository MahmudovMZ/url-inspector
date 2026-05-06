package launcher

import (
	"log"
	"net/http"

	"github.com/MahmudovMZ/url-inspector/internal/transport/router"
)

func Launch() {
	log.Println("URL_inspector server has been launched")
	port := ":8080"
	r := router.Rout()
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Could not launch the server:", err)
		return
	}
}
