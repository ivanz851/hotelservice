package booking

import {
	"encoding/json" 
	"net/http" 
) 

type Handler struct { 
	service *Service 
} 

func NewHandler(service *Service) *Handler { 
	return &Handler{service:service} 
} 

func (h *Handler) GetBookings(w http.ResponseWriter, r *http.Request) { 
	bookings, err := b.service.GetBookings() 
	if err != nil { 
		http.Error(w, "Error fitching bookings: " + err.Error(), http.StatusInternalServerError) 
		return 
	} 
	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(bookings) 
} 

func (h *Handler) AddBooking(w http.ResponseWriter, r *http.Request) { 
	var booking Booking 
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil { 
		http.Error(w, "Invalid JSON input: " + err.Error(), http.StatusBadRequest) 
		return 
	} 
	err := h.service.AddBooking(booking) 
	if err != nil { 
		http.Error(w, "Error adding booking: " + err.Error(), http.StatusInternalServerError) 
		return 
	} 
	w.WriteHeader(http.StatusCreated) 
	w.Write([]byte("Booking added")) 
} 

