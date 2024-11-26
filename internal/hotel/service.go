package hotel

type Service struct { 
	storage *Storage 
} 

func NewService(storage *Storage) *Service { 
	return &Service{storage:storage} 
} 

func (s *Service) GetHotels() ([]Hotel, error) { 
	return s.storage.GetHotels() 
} 

func (s *Service) AddHotel(hotel Hotel) error { 
	return s.storage.AddHotel(hotel) 
} 

