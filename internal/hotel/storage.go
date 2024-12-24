package hotel

import (
	"database/sql"
	"errors"
	"log"

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
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

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

func (s *Storage) GetHotelById(id int32) (models.Hotel, error) {
	var hotel models.Hotel
	err := s.db.QueryRow("SELECT id, name, address, price_per_night FROM hotels WHERE id = $1", id).
		Scan(&hotel.ID, &hotel.Name, &hotel.Address, &hotel.PricePerNight)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Hotel not found with ID: %d", id)
			return hotel, nil
		}
		log.Printf("Error fetching hotel by ID: %v", err)
		return hotel, err
	}
	return hotel, nil
}
