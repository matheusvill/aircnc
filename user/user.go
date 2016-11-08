package user

import (
	"errors"
	"regexp"
	"strings"

	"github.com/matheusvill/aircnc/storage"
)

type User struct {
	Email    string
	Password string
}

func Create(email, password string) error {
	s := storage.New()

	if checkEmail(email) {
		s.InsertUser(email, password)
		return nil
	}

	return errors.New("Invalid email!")
}

func Update(id int, email, password string) error {
	s := storage.New()

	if checkEmail(email) {
		s.UpdateUser(id, email, password)
		return nil
	}

	return errors.New("Invalid email!")
}

func Get(id int) User {
	s := storage.New()
	userMap := s.GetUser(id)

	user := User{}
	if userMap != nil {
		user.Email = userMap["email"]
		user.Password = userMap["password"]
	}

	return user
}

func checkEmail(email string) bool {
	email = strings.ToUpper(email)
	match, _ := regexp.MatchString(`^[A-Z0-9_.-]*[@][A-Z0-9]*[.][A-Z]{2,9}(\.[A-Z]{2,4})?(\.[A-Z]{2,4})?$`, email)
	if match {
		return true
	}
	return false
}
