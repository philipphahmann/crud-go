package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	rest_err "github.com/philipphahmann/crud-go/src/configuration/rest_error"
	"github.com/philipphahmann/crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userResquest request.UserRequest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, erro=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userResquest)
}
