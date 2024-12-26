package hotel

import (
	"database/sql"
	"fmt"
	"hotelservice/internal/models"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(conn string) *Storage {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic("Connection failed: " + err.Error())
	}
	if err := db.Ping(); err != nil {
		panic("Unable to connect to the database: " + err.Error())
	}
	return &Storage{db: db}
}

func (s *Storage) GetHotels() ([]models.Hotel, error) {
	rows, err := s.db.Query("SELECT id, name, address, email FROM hotels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []models.Hotel
	for rows.Next() {
		var hotel models.Hotel
		if err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Address, &hotel.Email); err != nil {
			if err != nil {
				return []models.Hotel{}, fmt.Errorf("no hotels found")
			}
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	log.Printf(hotels[0].Name)
	return hotels, nil
}

func (s *Storage) GetHotel(hotel_id int) (models.Hotel, error) {
	rows := s.db.QueryRow("SELECT id, name, address, email FROM hotels where id = $1", hotel_id)

	var hotel models.Hotel

	if err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Address, &hotel.Email); err != nil {
		if err != nil {
			return models.Hotel{}, fmt.Errorf("no hotels found")
		}
		return models.Hotel{}, err
	}
	return hotel, nil
}

func (s *Storage) AddHotel(hotel models.Hotel) error {
	_, err := s.db.Exec(
		`INSERT INTO hotels (
			name,
			address,
			price_per_night,
			email
		) VALUES ($1, $2, $3, $4)`,
		hotel.Name,
		hotel.Address,
		hotel.PricePerNight,
		hotel.Email,
	)
	return err
}
