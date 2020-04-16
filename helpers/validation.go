package helpers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"gopkg.in/go-playground/validator.v9"
)

func ValidationResponse(err error, c echo.Context)  {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required", err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s is not valid email", err.Field())
			case "gte":
				report.Message = fmt.Sprintf("%s harus lebih gede dari %s", err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s harus lebih kecil dari %s", err.Field(), err.Param())
			}
			break
		}
	}

	c.Logger().Error(report)
	c.JSON(report.Code, report)
}