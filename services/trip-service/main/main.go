package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"ride-sharing/services/trip-service/internal/api/grpc"
	httpApi "ride-sharing/services/trip-service/internal/api/http"
	"ride-sharing/services/trip-service/internal/core"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"syscall"
	"time"

	grpcserver "google.golang.org/grpc"
)

var GrpcPort = ":9093"
var httpPort = ":8083"

func startGRPCServer(ctx *context.Context, svc core.TripService) {
	log.Println("Starting the grpc server")
	lis, err := net.Listen("tcp", GrpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// starting the grpc server
	grpcServer := grpcserver.NewServer()
	grpc.NewGRPCHandler(grpcServer, svc)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve  :  %v", err)
		}
	}()

	log.Printf("Started grpc server on port %s", lis.Addr().String())
	// wait for the shutdown signal
	<-(*ctx).Done()
	log.Println("Shutting down the server...")
	grpcServer.GracefulStop()
}

func statHTTPserver(ctx *context.Context) {
	log.Println("Starting the http server")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /trip/preview", httpApi.HandleTripPreview)
	server := &http.Server{
		Addr:    httpPort,
		Handler: mux,
	}

	go func() {
		log.Printf("Server listening on %s", httpPort)
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Http server failed  %v", err)
		}
	}()

	// waiting for the shutdown signal
	<-(*ctx).Done()
	ctx2, _ := context.WithTimeout(context.Background(), 5*time.Second)
	log.Printf("Shutting down http server gracefully ")
	if err := server.Shutdown(ctx2); err != nil {
		log.Printf("could not shutdown server gracefully : %v", err)
		server.Close()
	}

}

func main() {

	inmemRepo := repository.NewInmemRepository()
	svc := core.NewService(inmemRepo)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdownSignal := make(chan os.Signal, 1)

	// this listen to os signal and relay that sigterm signal to sigCh
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	go startGRPCServer(&ctx, svc)
	go statHTTPserver(&ctx)

	<-shutdownSignal
	log.Printf("grpc and http servers forcefully shutting down")
	cancel()
}
