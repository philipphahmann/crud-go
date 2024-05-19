package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/philipphahmann/crud-go/src/configuration/validation"
	"github.com/philipphahmann/crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userResquest request.UserRequest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userResquest)
}
