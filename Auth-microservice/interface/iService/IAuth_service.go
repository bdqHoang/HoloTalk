package IService

import (
	"Auth-microservice/models"

	"github.com/bdqHoang/HoloTalk/shared/dto"
)

type IAuthService interface {
	Register(userReq dto.RegisterRequest) (*models.User, error)
	Login(loginReq dto.LoginRequest) (accessToen, refreshToken string, err error)
}
