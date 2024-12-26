package booking

import (
	"context"
	"encoding/json"
	"hotelservice/internal/models"
	pb "hotelservice/proto/booking"
	"log"
	"net/http"
)

type Handler struct {
	client pb.BookingServiceClient
}

func NewHandler(client pb.BookingServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.client.GetBookings(context.Background(), &pb.GetBookingsRequest{})
	if err != nil {
		http.Error(w, "Error fitching bookings: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Bookings: %v", bookings)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func (h *Handler) AddBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	ctx := context.Background()
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid JSON input: "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err := h.client.CreateBooking(ctx, &pb.CreateBookingRequest{HotelId: int32(booking.HotelID), ClientId: int32(booking.ClientID), Email: booking.Email})
	if err != nil {
		http.Error(w, "Error adding booking: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Booking added"))
}
