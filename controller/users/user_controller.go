package user

import (
	"fmt"
	"github.com/maik101010/oauthCourseGoLibrary/oauth"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/rest_errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maik101010/proyectCourseUsersApi/domain/user"
	service "github.com/maik101010/proyectCourseUsersApi/services"
)

const (
	//UserIDParam param solicitud service
	UserIDParam = "user_id"
)

func TestServiceInterface() {

}

//getUserById funtion get user by id
func getUserByID(id string) (int64, rest_errors.RestError) {
	userID, userErr := strconv.ParseInt(id, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("Invalid user id, give a number")
	}
	return userID, nil
}

//Create method
func Create(c *gin.Context) {
	var user user.User
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := rest_errors.NewBadRequestError("Invalidad json body")
		c.JSON(restError.Status(), restError)
		fmt.Println("Error: ", err.Error())
		return
	}
	result, saveErr := service.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	fmt.Println(user)
	// fmt.Println(string(bytes))
	// fmt.Println(err)
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil { //converter to user
	// 	fmt.Println("Error: ", err.Error())
	// 	return
	// }
}

//Update method
func Update(c *gin.Context) {
	userID, idErr := getUserByID(c.Param(UserIDParam))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := rest_errors.NewBadRequestError("Invalidad json body")
		c.JSON(restError.Status(), restError)
		return
	}
	user.ID = userID
	isPartial := c.Request.Method == http.MethodPatch
	result, updateErr := service.UsersService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

//Delete method
func Delete(c *gin.Context) {
	userID, idErr := getUserByID(c.Param(UserIDParam))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	getErr := service.UsersService.DeleteUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "delete"})
}

//Get method
func Get(c *gin.Context) {
	if err:= oauth.AuthenticateRequest(c.Request); err!=nil {
		c.JSON(err.Status, err)
		return
	}

	userID, idErr := getUserByID(c.Param(UserIDParam))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	user, getErr := service.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	if oauth.GetCallerId(c.Request)==user.ID {
		c.JSON(http.StatusOK, user.Marshal(false))
		return
	}
	c.JSON(http.StatusOK, user.Marshal(oauth.IsPublic(c.Request)))
}

//Search method controller find by status
func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := service.UsersService.Search(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))

}
func Login(c *gin.Context)  {
	var request user.LoginRequest
	if err:=c.ShouldBindJSON(&request); err!=nil {
		restErr:= rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr.Error)
		return
	}
	user, err := service.UsersService.LoginUser(request)
	if err!=nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// Truncate table users
func Truncate(c *gin.Context) {
	err := service.UsersService.TruncateUsers()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.String(http.StatusOK, "TruncateUsers")

}

//SearchUser method
// func SearchUser(c *gin.Context){
// 	c.String(http.StatusNotImplemented, "Implement me!")
// }
