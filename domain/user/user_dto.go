package user

import (
	"strings"
	"github.com/maik101010/proyectCourseUsersApi/utils/errors"
)

//User struct
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

//Users array
type Users []User

//Validate parameters user for struct
func (user *User) Validate() *errors.RestError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalidad email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("Invalidad password")
	}
	return nil
}

//Validate parameters user
// func Validate(user *User) *errors.RestError {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return nil, errors.NewBadRequestError("Invalidad email address")
// 	}
// 	return nil
// }
