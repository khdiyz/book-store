package handlerfunc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	userIdCtx    = "userId"
	userRoleCtx  = "userRole"
	paramInvalid = "%s param is invalid"
	queryInvalid = "%s query is invalid"
	ParseDate    = "2006-01-02T05:00:00.000Z"
	FormatDate   = "2006-01-02T05:00:00.000Z"
	orderAsc     = "asc"
	orderDesc    = "desc"
)

func GetNullStringQuery(c echo.Context, query string) string {
	param := c.QueryParam(query)
	param = strings.TrimSpace(param)
	return param
}

func GetIntParam(c echo.Context, query string) (int, error) {
	queryData := c.Param(query)
	if queryData == "" {
		err := fmt.Sprintf(queryInvalid, queryData)
		return 0, errors.New(err)
	}
	queryInt, err := strconv.Atoi(queryData)
	if err != nil {
		err := fmt.Sprintf(queryInvalid, queryData)
		return 0, errors.New(err)
	}

	return queryInt, nil
}
