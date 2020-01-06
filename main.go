package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/chadgrant/go-http-infra/infra"
	"github.com/chadgrant/go-http-infra/gorilla"
	"github.com/chadgrant/service-status/api/handlers"
	"github.com/chadgrant/service-status/api/repository/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	host := *flag.String("host", infra.GetEnvVar("SVC_HOST", "0.0.0.0"), "default binding 0.0.0.0")
	port := *flag.Int("port", infra.GetEnvVarInt("SVC_PORT", 8080), "default port 8080")

	mhost := *flag.String("mhost", infra.GetEnvVar("MYSQL_HOST", "localhost"), "default mysql host localhost")
	mport := *flag.Int("mport", infra.GetEnvVarInt("MYSQL_PORT", 3306), "default port 8080")
	muser := *flag.String("muser", infra.GetEnvVar("MYSQL_USER", "root"), "default user root")
	mpassword := *flag.String("mpassword", infra.GetEnvVar("MYSQL_PASSWORD", ""), "")
	mdb := *flag.String("mdb", infra.GetEnvVar("MYSQL_DATABASE", "service_status"), "default mysql database service_status")
	flag.Parse()

	eh := handlers.NewEnvironmentHandler(mysql.NewEnvironmentRepository(mhost, mport, muser, mpassword, mdb))
	sh := handlers.NewServiceHandler(mysql.NewServiceRepository(mhost, mport, muser, mpassword, mdb))
	dh := handlers.NewDeploymentHandler(mysql.NewDeploymentRepository(mhost, mport, muser, mpassword, mdb))

	r := mux.NewRouter()
	r.StrictSlash(true)
	gorilla.Handle(r)
	r.Use(infra.Recovery)

	r.HandleFunc("/environments/", eh.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/environments/", eh.Add).Methods(http.MethodPost)
	r.HandleFunc("/environment/{friendly}", eh.Update).Methods(http.MethodPut)

	r.HandleFunc("/environment/{environment}/services", sh.GetForEnvironment).Methods(http.MethodGet)

	r.HandleFunc("/environment/{environment}/deployments", dh.GetForEnvironmentPaged).Methods(http.MethodGet)
	r.HandleFunc("/environment/{environment}/service/{service}/deployments", dh.GetForServicePaged).Methods(http.MethodGet)

	r.HandleFunc("/services/", sh.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/services/", sh.Add).Methods(http.MethodPost)
	r.HandleFunc("/service/{friendly}", sh.Update).Methods(http.MethodPut)

	r.HandleFunc("/deployments", dh.GetPaged).Methods(http.MethodGet)

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
