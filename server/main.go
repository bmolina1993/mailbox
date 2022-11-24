package main

import (
	"fmt"
	"net/http"

	service "github.com/bmolina1993/mailbox/src/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%v", data)))
	})
	http.ListenAndServe(":8080", r)

}
