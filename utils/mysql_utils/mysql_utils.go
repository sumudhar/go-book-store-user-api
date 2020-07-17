package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sumudhar/go-book-store-user-api/utils/errors"
)

const(
	errorNoRows= "no rows in result set"

)

func ParseError(err error) *errors.RestErr{

	sqlErr,ok:= err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(),errorNoRows){
			return errors.NewBadRequestError("record not found with the given id")
		}
	    return  errors.NewInternalServerError("error parsing database response")
	} 
	switch sqlErr.Number{
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")

}
