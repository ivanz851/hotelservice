package booking

import (
	"context"
	pb "hotelservice/proto/booking"
	"log"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	pb.UnimplementedBookingServiceServer
	DB *sqlx.DB
}

func (s *Server) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	log.Printf("Creating booking for hotel_id: %d and client_id: %d", req.HotelId, req.ClientId)

	return &pb.CreateBookingResponse{
		Message:   "Booking created successfully",
		BookingId: 1,
	}, nil
}

func (s *Server) GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.GetBookingResponse, error) {
	log.Printf("Getting booking information for booking_id: %d", req.BookingId)

	return &pb.GetBookingResponse{
		BookingId: req.BookingId,
		HotelId:   1,
		ClientId:  123,
		Status:    "confirmed",
	}, nil
}
