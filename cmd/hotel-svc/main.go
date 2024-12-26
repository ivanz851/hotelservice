package main

import (
	"log"
	"net"
	"net/http"

	"hotelservice/internal/hotel"
	pb "hotelservice/proto/hotel"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	storage := hotel.NewStorage("postgres://user:password@postgres-hotel:5432/hotel_db?sslmode=disable")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	hotelService := &hotel.Server{Storage: storage}

	pb.RegisterHotelServiceServer(grpcServer, hotelService)

	r := mux.NewRouter()

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	hotelClient := pb.NewHotelServiceClient(conn)
	handler := hotel.NewHandler(hotelClient)

	r.HandleFunc("/hotels", handler.GetHotels).Methods("GET")
	r.HandleFunc("/hotels", handler.AddHotel).Methods("POST")

	go func() {
		log.Println("HTTP server started on port 8081")
		if err := http.ListenAndServe(":8081", r); err != nil {
			log.Fatalf("Failed to serve HTTP: %v", err)
		}
	}()

	log.Println("Hotel Service started on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
