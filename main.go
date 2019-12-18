package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/chadgrant/go/http/infra"
	"github.com/chadgrant/go/http/infra/gorilla"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	host := *flag.String("host", infra.GetEnvVar("SVC_HOST", "0.0.0.0"), "default binding 0.0.0.0")
	port := *flag.Int("port", infra.GetEnvVarInt("SVC_PORT", 8080), "default port 8080")
	flag.Parse()

	r := mux.NewRouter()
	gorilla.Handle(r)
	r.Use(infra.Recovery)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./docs/")))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		ExposedHeaders:   []string{"Location"},
		MaxAge:           86400,
	})

	log.Printf("Started, serving at %s:%d\n", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), c.Handler(r)))
}
