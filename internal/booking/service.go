package booking 

type Service struct { 
	storage *Storage 
} 

func NewService(storage *Storage) *Service {
	return &Service{storage:storage} 
} 

func (s *Service) GetBookings() ([]Booking, error) { 
	return s.storage.GetBookings() 
} 

func (s *Service) AddBooking(booking Booking) error { 
	return s.storage.AddBooking(booking) 
} 

