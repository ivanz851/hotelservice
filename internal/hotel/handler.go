package hotel 

import ( 
	"encoding/json" 
	"net/http" 
) 

type Handler struct { 
	service *Service 
} 

func NewHandler(service *Service) *Handler { 
	return &Handler{service:service} 
} 

func (h *Handler) GetHotels(w http.ResponseWriter, r *http.Request) { 
	hotels, err := h.service.GetHotels()
	if err != nil { 
		http.Error(w, "Error fetching hotels: " + err.Error(), http.StatusInternalServerError)
		return 
	} 
	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(hotels) 
} 

func (h *Handler) AddHotel(w http.ResponseWriter, r *http.Request) { 
	var hotel Hotel
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil { 
		http.Error(w, "Invalid JSON input: " + err.Error(), http.StatusBadRequest) 
		return 
	} 
	err := h.service.AddHotel(hotel) 
	if err != nil { 
		http.Error(w, "Error occured with hotel adding: " + err.Error(), http.StatusInternalServerError) 
		return 
	} 
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Hotel added")) 
} 

