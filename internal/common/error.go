package common

import "github.com/labstack/echo/v4"

type AppError struct {
	Code    int    `json:"-"` // HTTP Status Code
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}
func Error(c echo.Context, err error) error {
	if appErr, ok := err.(*AppError); ok {
		return c.JSON(appErr.Code, appErr)
	}
	return c.JSON(500, err.Error())
}
