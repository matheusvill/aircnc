package user

import "fmt"

type User struct {
	// repository repository.Repository

}

func (self *Sources) Create() {
	fmt.Println("a")
}

func NewUser() *User {
	log := log.NewLogger("user")
	// repository := repository.NewPostgres()
	self := Sources{
	// repository: repository,
	}

	return &self
}
