package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	pb "github.com/phucvin/project-teneng/services/servicerouter/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serviceRouterAddr := os.Getenv("SERVICE_ROUTER_ADDRESS")
	if serviceRouterAddr == "" {
		log.Fatalf("require SERVICE_ROUTER_ADDRESS")
	}
	serviceRouterConn, err := grpc.Dial(serviceRouterAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to ServiceRouter: %v", err)
	}
	defer serviceRouterConn.Close()
	serviceRouter := pb.NewServiceRouterClient(serviceRouterConn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := serviceRouter.Invoke(ctx, &pb.InvokeRequest{Description: "From: Sample To: Something"})
		if err != nil {
			fmt.Fprintf(w, "INTERNAL_ERROR: %v", err)
			return
		}

		fmt.Fprintf(w, "Hello, this is a sample app, serving from %s. Response from ServiceRouter: %v", os.Getenv("FLY_REGION"), res)
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
