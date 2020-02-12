package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chadgrant/go-http-infra/infra"
	"github.com/chadgrant/go-http-infra/infra/health"
	"github.com/chadgrant/go-http-infra/infra/schema"
	"github.com/chadgrant/service-status/api/handlers"
	"github.com/chadgrant/service-status/api/repository/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

func main() {
	host := *flag.String("host", infra.GetEnvVar("SVC_HOST", "0.0.0.0"), "default binding 0.0.0.0")
	port := *flag.Int("port", infra.GetEnvVarInt("SVC_PORT", 8080), "default port 8080")

	mhost := *flag.String("mhost", infra.GetEnvVar("MYSQL_HOST", "localhost"), "default mysql host localhost")
	mport := *flag.Int("mport", infra.GetEnvVarInt("MYSQL_PORT", 3306), "default port 8080")
	muser := *flag.String("muser", infra.GetEnvVar("MYSQL_USER", "docker"), "default user root")
	mpassword := *flag.String("mpassword", infra.GetEnvVar("MYSQL_PASSWORD", "password"), "")
	mdb := *flag.String("mdb", infra.GetEnvVar("MYSQL_DATABASE", "service_status"), "default mysql database service_status")
	flag.Parse()

	eh := handlers.NewEnvironmentHandler(mysql.NewEnvironmentRepository(mhost, mport, muser, mpassword, mdb))
	sh := handlers.NewServiceHandler(mysql.NewServiceRepository(mhost, mport, muser, mpassword, mdb))
	dh := handlers.NewDeploymentHandler(mysql.NewDeploymentRepository(mhost, mport, muser, mpassword, mdb))

	r := mux.NewRouter().StrictSlash(false)
	r.Use(infra.Recovery)

	gorillaW := func(s string, w http.HandlerFunc) {
		r.HandleFunc(s, w)
	}
	checker := health.NewHealthChecker()
	schemas := schema.NewRegistry()
	if err := infra.RegisterInfraHandlers(gorillaW, checker, schemas); err != nil {
		panic(err)
	}

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true&interpolateParams=true", muser, mpassword, mhost, mport, mdb))
	if err != nil {
		panic(fmt.Errorf("could not connect to database %v", err))
	}
	defer db.Close()

	checker.AddReadiness("mysql", time.Second*10, health.DatabasePingCheck(db, time.Second*3))
	checker.AddReadiness("google tcp connection", time.Second*10, health.TCPDialCheck("google.com:80", 3*time.Second))
	checker.AddReadiness("http get", time.Second*10, health.HTTPGetCheck("https://golang.org", 3*time.Second))
	checker.AddReadiness("dns lookup", time.Second*10, health.DNSResolveCheck("google.com", 3*time.Second))

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
