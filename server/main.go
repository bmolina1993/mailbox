package main

import (
	"fmt"
	"net/http"

	service "github.com/bmolina1993/mailbox/src/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// obtiene data de zinc
	data, err := service.FetchData()
	if err != nil {
		fmt.Println(err)
	}

	// habilita servidor con data recolectada
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//habilita cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%v", data)))
	})
	http.ListenAndServe(":8080", r)

}
