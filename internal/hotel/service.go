package hotel

import (
	"context"
	"hotelservice/internal/models"
	pb "hotelservice/proto/hotel"
	"log"
)

type Server struct {
	pb.UnimplementedHotelServiceServer
	storage *Storage
}

func (s *Server) CreateHotel(ctx context.Context, req *pb.CreateHotelRequest) (*pb.CreateHotelResponse, error) {
	log.Printf("Creating hotel: %s at %s with price: %.2f", req.Name, req.Address, req.PricePerNight)
	s.storage.AddHotel(models.Hotel{Name: req.Name, Address: req.Address})
	return &pb.CreateHotelResponse{
		Message: "Hotel created successfully",
		HotelId: 1,
	}, nil
}

func (s *Server) GetHotel(ctx context.Context, req *pb.GetHotelRequest) (*pb.GetHotelResponse, error) {
	log.Printf("Getting hotel information for hotel_id:d", req.HotelId)

	hotel, err := s.storage.GetHotelById(req.HotelId)
	if err != nil {
		log.Printf("Error fetching hotel: %v", err)
		return nil, err
	}

	return &pb.GetHotelResponse{
		HotelId:       hotel.ID,
		Name:          hotel.Name,
		Address:       hotel.Address,
		PricePerNight: hotel.PricePerNight,
	}, nil
}
