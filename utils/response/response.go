package response

import (
	"github.com/labstack/echo/v4"
)

// HandleResponse handles sending a JSON response with the provided status, error, data, and pagination.
func HandleResponse(ctx echo.Context, status Status, err error, data interface{}, pagination interface{}) error {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	response := ResponseListModel{
		Status:       status.Status,
		Code:         status.Code,
		Description:  status.Description,
		Pagination:   pagination,
		SnapData:     data,
		ErrorMessage: errorMessage,
	}
	return ctx.JSON(status.Code, response)
}

// ResponseListModel defines the structure of the JSON response.
type ResponseListModel struct {
	Status       string      `json:"status,omitempty"`
	Code         int         `json:"code,omitempty"`
	Description  string      `json:"description,omitempty"`
	SnapData     interface{} `json:"snapData,omitempty"`
	Pagination   interface{} `json:"pagination,omitempty"`
	ErrorMessage string      `json:"error,omitempty"`
}

// ResponseModel defines the structure of a simplified JSON response.
type ResponseModel struct {
	Status       string      `json:"status,omitempty"`
	Code         int         `json:"code,omitempty"`
	Description  string      `json:"description,omitempty"`
	SnapData     interface{} `json:"snapData,omitempty"`
	ErrorMessage string      `json:"error,omitempty"`
}
