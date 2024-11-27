package hotel

import (
	"context"
	pb "hotelservice/proto/hotel"
	"log"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	pb.UnimplementedHotelServiceServer
	DB *sqlx.DB
}

func (s *Server) CreateHotel(ctx context.Context, req *pb.CreateHotelRequest) (*pb.CreateHotelResponse, error) {
	log.Printf("Creating hotel: %s at %s with price: %.2f", req.Name, req.Address, req.PricePerNight)

	return &pb.CreateHotelResponse{
		Message: "Hotel created successfully",
		HotelId: 1,
	}, nil
}

func (s *Server) GetHotel(ctx context.Context, req *pb.GetHotelRequest) (*pb.GetHotelResponse, error) {
	log.Printf("Getting hotel information for hotel_id:d", req.HotelId)

	return &pb.GetHotelResponse{
		HotelId:       req.HotelId,
		Name:          "Sample Hotel",
		Address:       "123 Sample St",
		PricePerNight: 100.0,
	}, nil
}
