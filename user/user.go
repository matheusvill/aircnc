package user

import (
	"errors"
	"regexp"
	"strings"

	"github.com/matheusvill/aircnc/storage"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Create(user User) error {
	if CheckEmail(user.Email) {
		storage.InsertUser(user.Email, user.Password)
		return nil
	}

	return errors.New("Invalid email!")
}

// func Update(id int, email, password string) error {
// 	if checkEmail(email) {
// 		storage.UpdateUser(id, email, password)
// 		return nil
// 	}

// 	return errors.New("Invalid email!")
// }

func Get(id int) User {
	userMap := storage.GetUser(id)

	user := User{}
	if userMap != nil {
		user.Email = userMap["email"].(string)
		user.Password = userMap["password"].(string)
	}

	return user
}

func CheckEmail(email string) bool {
	email = strings.ToUpper(email)
	match, _ := regexp.MatchString(`^[A-Z0-9_.-]*[@][A-Z0-9]*[.][A-Z]{2,9}(\.[A-Z]{2,4})?(\.[A-Z]{2,4})?$`, email)
	if match {
		return true
	}
	return false
}
