package mysql_utils

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/rest_errors"
	"strings"
)

const (
	NoRowResultSet = "no rows in result set"
)

//ParseError function converter error
func ParseError(err error) *rest_errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), NoRowResultSet) {
			return rest_errors.NewNotFoundError("No record matching given id")
		}
		return rest_errors.NewInternalServerError("Error parsing database response", errors.New("database error"))
	}
	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("Invalid data")
	}
	return rest_errors.NewInternalServerError("Error processing request", errors.New("database error"))
}
