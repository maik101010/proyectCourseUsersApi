package user

import (
	"errors"
	"fmt"
	"github.com/maik101010/proyectCourseUsersApi/datasource/mysql/user_db_client"
	"github.com/maik101010/proyectCourseUsersApi/logger"
	"github.com/maik101010/proyectCourseUsersApi/utils/mysql_utils"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/rest_errors"
	"strings"
)

const (
	//QueryInsertUser constan insert user
	QueryInsertUser = "INSERT INTO user(first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	//QuerySelecttUser constan select user by id
	QuerySelecttUser = "SELECT id, first_name, last_name, email, date_created, status FROM user WHERE id=?;"
	//QueryUpdateUser constan update user
	QueryUpdateUser = "UPDATE user SET first_name =?, last_name=?, email=? WHERE id=?;"
	//QueryDeleteUser constan update user
	QueryDeleteUser = "DELETE FROM user WHERE id=?;"
	//FindUserByStatus constan update user
	FindUserByStatus = "SELECT id, first_name, last_name, email, date_created FROM user WHERE status=?;"
	//FindByEmailAndPassword constan find user by email and password
	FindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM user WHERE email=? AND password =? AND status=?;"
	//QueryTruncateUser constan update user
	QueryTruncateUser = "TRUNCATE TABLE user;"
	indexUniqueEmail  = "email"
)

//Get user by id
func (user *User) Get() *rest_errors.RestError {
	statement, err := user_db_client.ClientDb.Prepare(QuerySelecttUser)
	if err != nil {
		logger.Error("Error trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("error when trying to get user error", errors.New("database error"))
	}
	defer statement.Close()
	result := statement.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("Error trying to get user by id", err)
		return rest_errors.NewInternalServerError("error when trying to get user error", errors.New("database error"))
	}
	return nil
}

//Save user
func (user *User) Save() * rest_errors.RestError {
	statement, err := user_db_client.ClientDb.Prepare(QueryInsertUser)
	if err != nil {
		logger.Error("Error trying to prepare save user statement", err)
		return rest_errors.NewInternalServerError("error when trying to get user error", errors.New("database error"))
	}
	defer statement.Close() //cuando termine lo demás, cierra la conexión
	inserResult, saveError := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveError != nil {
		logger.Error("Error exec save user", err)
		return rest_errors.NewInternalServerError("error when trying to get user error", errors.New("database error"))
		// fmt.Println(sqlError.Number)
		// fmt.Println(sqlError.Message)
	}
	// result, err := user_db_client.Exec(QueryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	idUser, err := inserResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to get last insert id after creating a new user save user", err)
		return rest_errors.NewInternalServerError("error when trying to get user error", errors.New("database error"))
		user.ID = idUser
		return nil
	}
	return nil
}
	//Update function update user database
	func (user *User) Update() *rest_errors.RestError {
		statement, err := user_db_client.ClientDb.Prepare(QueryUpdateUser)
		if err != nil {
		logger.Error("Error when trying to update save statement", err)
			return rest_errors.NewInternalServerError("error when trying to update user error", errors.New("database error"))
	}
		defer statement.Close()
		_, err = statement.Exec(user.FirstName, user.LastName, user.Email, user.ID)
		if err != nil {
		logger.Error("Error when trying to update user exec", err)
			return rest_errors.NewInternalServerError("error when trying to update user error", errors.New("database error"))
	}
		return nil
	}

	//Delete function update user database
	func (user *User) Delete() *rest_errors.RestError {
		statement, err := user_db_client.ClientDb.Prepare(QueryDeleteUser)
		if err != nil {
		logger.Error("Error when prepare to delete user", err)
			return rest_errors.NewInternalServerError("error when trying to delete user error", errors.New("database error"))
	}
		defer statement.Close()
		_, err = statement.Exec(user.ID)
		if err != nil {
		logger.Error("Error when trying delete user", err)
			return rest_errors.NewInternalServerError("error when trying to delete user error", errors.New("database error"))
	}
		return nil
	}

	//FindByStatus find users by status
	func (user *User) FindByStatus(status string) ([]User, *rest_errors.RestError) {
		statement, err := user_db_client.ClientDb.Prepare(FindUserByStatus)
		if err != nil {
			logger.Error("Error when prepare to find user by status", err)
			return nil, rest_errors.NewInternalServerError("error when trying to find user error", errors.New("database error"))
		}
		defer statement.Close()
		rows, err := statement.Query(status)
		defer rows.Close()
		if err != nil {
			logger.Error("Error when exect scan user to find user by status", err)
			return nil, rest_errors.NewInternalServerError("error when trying to find user error", errors.New("database error"))
		}
		results := make([]User, 0)
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
				return nil, mysql_utils.ParseError(err)
			}
			results = append(results, user)
		}
		if len(results) == 0 {
			return nil, rest_errors.NewNotFoundError(fmt.Sprintf("No users matching status %s", status))
		}
		return results, nil
	}
//Get user by id
	func (user *User) FindByEmailAndPassword() *rest_errors.RestError {
		statement, err := user_db_client.ClientDb.Prepare(FindByEmailAndPassword)
		if err != nil {
			logger.Error("Error trying to prepare get user by email and password", err)
			return rest_errors.NewInternalServerError("error when trying to find user error", errors.New("database error"))
		}
		defer statement.Close()
		result := statement.QueryRow(user.Email, user.Password, StatusActive)
		if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			if strings.Contains(err.Error(), mysql_utils.NoRowResultSet) {
				return rest_errors.NewNotFoundError("invalid user credentials")
			}
			logger.Error("Error trying to get user by email and password", err)
			return rest_errors.NewInternalServerError("error when trying to get find error", errors.New("database error"))
		}
		return nil
	}

	// TruncateUsers function truncate
	func TruncateUsers() *rest_errors.RestError {
		statement, err := user_db_client.ClientDb.Prepare(QueryTruncateUser)
		if err != nil {
			return rest_errors.NewInternalServerError("error when trying to truncate user error", errors.New("database error"))
		}
		defer statement.Close()
		_, err = statement.Exec()
		if err != nil {
		return mysql_utils.ParseError(err)
	}
		return nil
	}
