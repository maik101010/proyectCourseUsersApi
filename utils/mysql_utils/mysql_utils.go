package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/maik101010/proyectCourseUsersApi/utils/errors"
)

const (
	NoRowResultSet = "no rows in result set"
)

//ParseError function converter error
func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), NoRowResultSet) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError("Error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid data")
	}
	return errors.NewInternalServerError("Error processing request")
}
