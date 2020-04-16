package service

import (
	"github.com/maik101010/proyectCourseUsersApi/domain/user"
	users "github.com/maik101010/proyectCourseUsersApi/domain/user"
	"github.com/maik101010/proyectCourseUsersApi/utils/date_utils"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/crypto_utils"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/rest_errors"
)

var (
	//UsersService Interface service user
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}
type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, rest_errors.RestError)
	UpdateUser(bool, users.User) (*users.User, rest_errors.RestError)
	GetUser(int64) (*users.User, rest_errors.RestError)
	DeleteUser(int64) rest_errors.RestError
	Search(string) (users.Users, rest_errors.RestError)
	LoginUser(request user.LoginRequest) (*users.User, rest_errors.RestError)
	TruncateUsers() rest_errors.RestError
}

//CreateUser service method
func (s *usersService) CreateUser(user users.User) (*users.User, rest_errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDataBaseFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUser service method
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, rest_errors.RestError) {
	current := &users.User{ID: user.ID}
	if err := current.Get(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

//GetUser get user by id
func (s *usersService) GetUser(userID int64) (*users.User, rest_errors.RestError) {
	result := &users.User{ID: userID}

	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

//DeleteUser delete user by id
func (s *usersService) DeleteUser(userID int64) rest_errors.RestError {
	result := &users.User{ID: userID}
	return result.Delete()
}

//Search find by status user
func (s *usersService) Search(status string) (users.Users, rest_errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s * usersService) LoginUser(request user.LoginRequest) (*users.User, rest_errors.RestError){
	dao := &users.User{
		Email:request.Email,
		Password:crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err !=nil{
		return nil, err
	}
	return dao, nil
}


//TruncateUsers truncate table user
func (s *usersService) TruncateUsers() rest_errors.RestError {
	return user.TruncateUsers()
}
