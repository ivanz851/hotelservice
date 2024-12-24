package hotel

import (
	"database/sql"

	_ "github.com/lib/pq"
	"hotelservice/internal/models"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(conn string) *Storage {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic("Connection failed: " + err.Error())
	}
	return &Storage{db: db}
}

func (s *Storage) GetHotels() ([]models.Hotel, error) {
	rows, err := s.db.Query("SELECT id, name, address, price_per_night FROM hotels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []models.Hotel
	for rows.Next() {
		var hotel models.Hotel
		if err := rows.Scan(&hotel.ID, &hotel.Name); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (s *Storage) AddHotel(hotel models.Hotel) error {
	_, err := s.db.Exec(
		`INSERT INTO hotels (
			id,
			name,
			address,
            price_per_night
		) VALUES ($1, $2, $3, $4)`,
		hotel.ID,
		hotel.Name,
		hotel.Address,
		hotel.PricePerNight,
	)
	return err
}
