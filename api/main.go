package main

import (
	"log"
	"net"
	"net/http"

	generated "github.com/bromivipo/marketplace/api/definitions"
	grpchandlers "github.com/bromivipo/marketplace/api/handlers/grpc"
	"github.com/bromivipo/marketplace/api/handlers/http"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

func main() {
	go func() {
		lis, err := net.Listen("tcp", "localhost:50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		generated.RegisterMarketplaceInternalServer(grpcServer, &grpchandlers.GrpcServer{})
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	go func() {
		router := chi.NewRouter()
		handler := generated.NewStrictHandler(httphandlers.NewServer(), nil)
		generated.HandlerFromMux(handler, router)
		log.Println("http server started on port 8080")
		if err := http.ListenAndServe("localhost:8080", router); err != nil {
			log.Fatalf("failed to serve http: %v", err)
		}
	} ()
	
	select {}
}