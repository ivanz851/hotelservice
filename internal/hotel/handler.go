package hotel

import (
	"context"
	"encoding/json"
	"hotelservice/internal/models"
	pb "hotelservice/proto/hotel"
	"net/http"
)

type Handler struct {
	client pb.HotelServiceClient
}

func NewHandler(client pb.HotelServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) GetHotels(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	hotels, err := h.client.GetHotels(ctx, &pb.GetHotelsRequest{})
	if err != nil {
		http.Error(w, "Error fetching hotels: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotels)
}

func (h *Handler) AddHotel(w http.ResponseWriter, r *http.Request) {
	var hotel models.Hotel
	ctx := context.Background()
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		http.Error(w, "Invalid JSON input: "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err := h.client.CreateHotel(ctx, &pb.CreateHotelRequest{Name: hotel.Name, Address: hotel.Address, PricePerNight: hotel.PricePerNight, Email: hotel.Email})
	if err != nil {
		http.Error(w, "Error occured with hotel adding: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Hotel added"))
}
