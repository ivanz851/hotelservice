package hotel

import (
	"context"
	"fmt"
	"hotelservice/internal/models"
	pb "hotelservice/proto/hotel"
	"log"
)

type Server struct {
	pb.UnimplementedHotelServiceServer
	Storage *Storage
}

func (s *Server) CreateHotel(ctx context.Context, req *pb.CreateHotelRequest) (*pb.CreateHotelResponse, error) {
	log.Printf("Creating hotel: %s at %s with price: %.2f", req.Name, req.Address, req.PricePerNight)
	s.Storage.AddHotel(models.Hotel{Name: req.Name, Address: req.Address, PricePerNight: req.PricePerNight, Email: req.Email})
	return &pb.CreateHotelResponse{
		Message: "Hotel created successfully",
		HotelId: 1,
	}, nil
}

func (s *Server) GetHotel(ctx context.Context, req *pb.GetHotelRequest) (*pb.GetHotelResponse, error) {
	log.Printf("Getting hotel information for hotel_id:d", req.HotelId)
	hotel, err := s.Storage.GetHotel(int(req.HotelId))
	if err != nil {
		log.Printf("error getting hotel information for hotel_id:", req.HotelId)
		return nil, fmt.Errorf("no hotels found")
	}

	var foundHotel *pb.Hotel

	foundHotel = &pb.Hotel{
		HotelId:       int64(hotel.ID),
		Name:          hotel.Name,
		Address:       hotel.Address,
		PricePerNight: float32(hotel.PricePerNight),
	}

	return &pb.GetHotelResponse{
		Hotel: foundHotel,
	}, nil
}

func (s *Server) GetHotels(ctx context.Context, req *pb.GetHotelsRequest) (*pb.GetHotelsResponse, error) {
	log.Printf("Getting hotel information for hotel_id:d")
	hotels, err := s.Storage.GetHotels()
	if err != nil {
		log.Printf("error getting hotel information for hotel_id:")
		return nil, fmt.Errorf("no hotels found")
	}

	var foundHotels []*pb.Hotel
	for _, hotel := range hotels {

		foundHotels = append(foundHotels, &pb.Hotel{
			HotelId:       int64(hotel.ID),
			Name:          hotel.Name,
			Address:       hotel.Address,
			PricePerNight: 100.00,
		})
	}

	return &pb.GetHotelsResponse{
		Hotels: foundHotels,
	}, nil
}
