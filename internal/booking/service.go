package booking

import (
	"context"
	"fmt"
	"hotelservice/internal/models"
	pb "hotelservice/proto/booking"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

type Server struct {
	pb.UnimplementedBookingServiceServer
	Storage *Storage
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func (s *Server) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	log.Printf("Creating booking for hotel_id: %d and client_id: %d", req.HotelId, req.ClientId)

	s.Storage.AddBooking(models.Booking{HotelID: int(req.HotelId), ClientID: int(req.ClientId), Email: req.Email})

	writer := newKafkaWriter("kafka:9092", "bookings")
	defer writer.Close()
	msg := kafka.Message{
		Key:   []byte("email"),
		Value: []byte(req.Email + " " + string(req.HotelId)),
	}
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("produced")
	}

	return &pb.CreateBookingResponse{
		Message:   "Booking created successfully",
		BookingId: int32(1),
	}, nil
}

func (s *Server) GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.GetBookingResponse, error) {
	log.Printf("Getting booking information for booking_id: %d", req.BookingId)

	booking, err := s.Storage.GetBooking(int(req.BookingId))
	if err != nil {
		log.Printf("Error fetching booking: %v", err)
		return nil, err
	}
	return &pb.GetBookingResponse{
		Booking: &pb.Booking{
			BookingId: int32(booking.ID),
			HotelId:   int32(booking.HotelID),
			ClientId:  int32(booking.ClientID),
			Status:    "confirmed",
			Email:     booking.Email,
		},
	}, nil
}

func (s *Server) GetBookings(ctx context.Context, req *pb.GetBookingsRequest) (*pb.GetBookingsResponse, error) {
	log.Println("Getting all bookings")

	bookings, err := s.Storage.GetBookings()
	if err != nil {
		log.Printf("Error fetching bookings: %v", err)
		return nil, fmt.Errorf("error fetching bookings")
	}

	var bookingList []*pb.Booking
	for _, booking := range bookings {
		bookingList = append(bookingList, &pb.Booking{
			BookingId: int32(booking.ID),
			HotelId:   int32(booking.HotelID),
			ClientId:  int32(booking.ClientID),
			Status:    booking.Status,
			Email:     booking.Email,
		})
	}

	return &pb.GetBookingsResponse{
		Bookings: bookingList,
	}, nil
}
