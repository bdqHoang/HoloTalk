package handler

import "Auth-microservice/Interface/IService"

type AuthHandler struct {
	authService IService.IAuthService
}

func NewAuthHandler(authService IService.IAuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type RegisterRequest struct{

}