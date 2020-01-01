package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gogo/gateway"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

<<<<<<< HEAD
	"github.com/chadgrant/go/http/infra"
	"github.com/chadgrant/service-status/api/generated"
	"github.com/chadgrant/service-status/api/insecure"
	"github.com/chadgrant/service-status/api/repository"
=======
	"github.com/chadgrant/go-http-infra/infra"
	"github.com/chadgrant/go-http-infra/gorilla"
	"github.com/chadgrant/service-status/api/handlers"
>>>>>>> master
	"github.com/chadgrant/service-status/api/repository/mysql"
	"github.com/chadgrant/service-status/api/server"
)

func main() {
	host := *flag.String("host", infra.GetEnvVar("SVC_HOST", "0.0.0.0"), "default binding 0.0.0.0")
	port := *flag.Int("port", infra.GetEnvVarInt("SVC_PORT", 8080), "default http port 8080")
	tlsPort := *flag.Int("tls-port", infra.GetEnvVarInt("SVC_TLS_PORT", 8443), "default https port 8443")
	grpcPort := *flag.Int("grpc-port", infra.GetEnvVarInt("SVC_GRPC_PORT", 8888), "default grpc port 8888")

	mhost := *flag.String("mhost", infra.GetEnvVar("MYSQL_HOST", "localhost"), "default mysql host localhost")
	mport := *flag.Int("mport", infra.GetEnvVarInt("MYSQL_PORT", 3306), "default port 8080")
	muser := *flag.String("muser", infra.GetEnvVar("MYSQL_USER", "docker"), "default user docker")
	mpassword := *flag.String("mpassword", infra.GetEnvVar("MYSQL_PASSWORD", "password"), "")
	mdb := *flag.String("mdb", infra.GetEnvVar("MYSQL_DATABASE", "service_status"), "default mysql database service_status")
	flag.Parse()

	repos := map[string]interface{}{
		"environment": mysql.NewEnvironmentRepository(mhost, mport, muser, mpassword, mdb),
		"service":     mysql.NewServiceRepository(mhost, mport, muser, mpassword, mdb),
		"deployment":  mysql.NewDeploymentRepository(mhost, mport, muser, mpassword, mdb),
	}

	if err := grpcserve(host, grpcPort, repos); err != nil {
		log.Fatalln("could not start grpc service", err)
		os.Exit(-1)
	}

	mux, err := grpcgateway(host, grpcPort)
	if err != nil {
		log.Fatalln("could not start grpc gateway", err)
		os.Exit(-2)
	}

	if err := httpserve(mux, host, port, tlsPort); err != nil {
		log.Fatalln("could not start http services", err)
		os.Exit(-3)
	}
}

func grpcserve(host string, port int, repos map[string]interface{}) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
		return err
	}
	s := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
		grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
	)

	generated.RegisterEnvironmentsServer(s, server.NewEnvironmentServer(repos["environment"].(repository.EnvironmentRepository)))
	generated.RegisterDeploymentsServer(s, server.NewDeploymentServer(repos["deployment"].(repository.DeploymentRepository)))
	generated.RegisterServicesServer(s, server.NewServiceServer(repos["service"].(repository.ServiceRepository)))

	go func() {
		log.Printf("Serving grpc on https://%s:%d", host, port)
		log.Fatal(s.Serve(lis))
	}()
	return nil
}

func grpcgateway(host string, port int) (*runtime.ServeMux, error) {
	// See https://github.com/grpc/grpc/blob/master/doc/naming.md
	// for gRPC naming standard information.
	addr := fmt.Sprintf("passthrough://localhost/%s:%d", host, port)
	conn, err := grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
		return nil, err
	}

	jsonpb := &gateway.JSONPb{
		EmitDefaults: false,
		Indent:       "  ",
		OrigName:     true,
	}
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
		// This is necessary to get error details properly
		// marshalled in unary requests.
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	if err := generated.RegisterEnvironmentsHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
		return mux, err
	}
	if err := generated.RegisterDeploymentsHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
		return mux, err
	}
	if err := generated.RegisterServicesHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
		return mux, err
	}

	return mux, nil
}

func httpserve(rmux *runtime.ServeMux, host string, port, tlsport int) error {
	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	infra.HandleMux(mux)
	mux.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	c := corsOptions()

	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", host, tlsport),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{insecure.Cert},
		},
		Handler: c.Handler(mux),
	}

	go func() {
		log.Printf("Serving https gateway on https://%s:%d", host, tlsport)
		log.Printf("Serving https docs on https://%s:%d/docs/swagger", host, tlsport)
		log.Fatalln(server.ListenAndServeTLS("", ""))
	}()

	log.Printf("Serving http gateway on http://%s:%d\n", host, port)
	log.Printf("Serving http docs on http://%s:%d/docs/swagger", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), c.Handler(mux)))

	return nil
}

func corsOptions() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		ExposedHeaders:   []string{"Location"},
		MaxAge:           86400,
	})
}
