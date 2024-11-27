package hotel

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

type Hotel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func NewStorage(conn string) *Storage {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic("Connection failed: " + err.Error())
	}
	return &Storage{db: db}
}

func (s *Storage) GetHotels() ([]Hotel, error) {
	rows, err := s.db.Query("SELECT id, name, city FROM hotels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []Hotel
	for rows.Next() {
		var hotel Hotel
		if err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.City); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (s *Storage) AddHotel(hotel Hotel) error {
	_, err := s.db.Exec("INSERT INTO hotels (name, city) VALUES ($1, $2)", hotel.Name, hotel.City)
	return err
}
