package booking

import (
	"github.com/matheusvill/aircnc/apartment"
	"github.com/matheusvill/aircnc/storage"
	"github.com/matheusvill/aircnc/user"
)

type Booking struct {
	User      int `json:"user"`
	Apartment int `json:"apartment"`
}

func Create(booking Booking) error {

	storage.InsertBooking(booking.User, booking.Apartment)

	return nil
}

func Get(id int) map[string]interface{} {
	bookingMap := storage.GetBooking(id)
	user := user.Get(bookingMap["user"].(int))
	apartment := apartment.Get(bookingMap["apartment"].(int))

	result := make(map[string]interface{})
	result["user"] = user
	result["apartment"] = apartment

	return result
}
