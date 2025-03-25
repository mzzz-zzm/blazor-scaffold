package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/mzzz-zzm/blazor-scaffold/svr/gapi"
	pb "github.com/mzzz-zzm/blazor-scaffold/svr/pb/greet"
)

func main() {
	// Create a normal gRPC server
	grpcServer := grpc.NewServer()

	pb.RegisterGreeterServer(grpcServer, &gapi.Server{})
	reflection.Register(grpcServer)

	// Create a gRPC-Web wrapper for the gRPC server
	wrappedGrpc := grpcweb.WrapServer(grpcServer,
		grpcweb.WithWebsockets(true),
	)

	// Configure CORS
	corsHandler := cors.New(cors.Options{
		// AllowedOrigins:   []string{"http://localhost:5001"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"grpc-status", "grpc-message"},
		AllowCredentials: true,
	})

	// Create an HTTP server that proxies gRPC-Web requests to gRPC
	httpServer := &http.Server{
		Addr: ":8080",
		Handler: corsHandler.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Log all headers
			log.Printf("Received request with headers:")
			for name, values := range r.Header {
				for _, value := range values {
					log.Printf("  %s: %s", name, value)
				}
			}
			log.Printf("Received %s request for %s", r.Method, r.URL.Path)
			if wrappedGrpc.IsGrpcWebRequest(r) || wrappedGrpc.IsAcceptableGrpcCorsRequest(r) {
				wrappedGrpc.ServeHTTP(w, r)
				return
			}
			// Handle regular HTTP requests (helpful for checking server status)
			if r.URL.Path == "/health" {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Server is healthy"))
				return
			}
			// Handle regular HTTP requests
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("gRPC-Web proxy server running"))
		})),
	}

	go func() {
		log.Println("Starting gRPC-Web proxy server on :8080...")
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start server: %v\n", err)
		}
		log.Println("Server stopped")
	}()
	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
