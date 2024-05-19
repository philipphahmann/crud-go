package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philipphahmann/crud-go/src/configuration/logger"
	"github.com/philipphahmann/crud-go/src/configuration/validation"
	"github.com/philipphahmann/crud-go/src/controller/model/request"
	"github.com/philipphahmann/crud-go/src/model"
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

	domain := model.NewUserDomain(
		userResquest.Email,
		userResquest.Password,
		userResquest.Name,
		userResquest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User was created successfully",
		zap.String("journey", "createUser"),
	)

	c.String(http.StatusOK, "")
}
