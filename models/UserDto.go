package models

import "fmt"

type UserDto struct {
	Name    string
	Age     int
	Citizen bool
}

func (userDto UserDto) getUserInfo() string {
	return fmt.Sprintf("User : %v", userDto)
}
