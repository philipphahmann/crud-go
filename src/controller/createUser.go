package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philipphahmann/crud-go/src/configuration/logger"
	"github.com/philipphahmann/crud-go/src/configuration/validation"
	"github.com/philipphahmann/crud-go/src/controller/model/request"
	"github.com/philipphahmann/crud-go/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userResquest request.UserRequest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userResquest.Email,
		Name:  userResquest.Name,
		Age:   userResquest.Age,
	}

	logger.Info("User was created successfully",
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, response)
}
