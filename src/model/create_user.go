package model

import (
	"github.com/philipphahmann/crud-go/src/configuration/logger"
	rest_err "github.com/philipphahmann/crud-go/src/configuration/rest_error"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	
	logger.Info("Init createUser model", zap.String("journey", "createdUser"))

	ud.EncryptPassword()

	return nil
}
