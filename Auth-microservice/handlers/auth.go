package handlers

import (
	"Auth-microservice/config"
	"Auth-microservice/db"
	"Auth-microservice/models"
	"Auth-microservice/service"
	"context"
	"time"

	"github.com/bdqHoang/HoloTalk/shared/dto"
	"github.com/bdqHoang/HoloTalk/shared/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var env = config.LoadEnv()

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
}

// handler to register
func (s *AuthServer)Register( ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {

	// convert proto request to dto
	userReq := dto.RegisterRequest{
		Email:           req.Email,
		Phone:           req.Phone,
		Name:            req.Name,
		DateOfBirth:     req.DateOfBirth,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
		Role:            uint(req.Role),
	}

	// validate register request
	err := service.ValidateRegisterRequest(userReq)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// hash password
	hashedPassword, err := service.HashPassord(userReq.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to hash password")
	}

	// create user
	user := models.User{
		Email:           userReq.Email,
		Phone:           userReq.Phone,
		Name:            userReq.Name,
		DateOfBirth:     userReq.DateOfBirth,
		PasswordHash:    string(hashedPassword),
		RoleID:          userReq.Role,
	}

	if err:= db.Db_context.Create(&user).Error; err != nil {
		return nil, status.Error(codes.Internal, "Failed to create user")
	}

	return &proto.RegisterResponse{Message: "User created successfully"}, nil
}

// handler to login
func (s *AuthServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.AuthResponse, error) {
	// convert proto request to dto
	userReq := dto.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	// validate login request
	err := service.ValidateLoginRequest(userReq)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// get user from db
	var user models.User
	if err := db.Db_context.Preload("Role").First(&user, "email = ?", userReq.Username, "phone = ?", userReq.Username).Error; err != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	// check password
	if err := service.CheckPassword(user.PasswordHash, userReq.Password); err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid password")
	}

	// generate token
	accessToken, err := service.GenerateAccessToken(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to generate access token")
	}

	refreshToken, err := service.GenerateRefreshToken(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to generate refresh token")
	}

	// update token in db
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	user.ExpiresIn = time.Now().Add(time.Duration(env.ACCESS_TOKEN_EXP))
	user.ExpiresAt = time.Now().Add(time.Duration(env.REFRESH_TOKEN_EXP))

	if err := db.Db_context.Save(&user).Error; err != nil {
		return nil, status.Error(codes.Internal, "Failed to save token in db")
	}

	return &proto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    user.ExpiresIn.Unix(),
		ExpiresAt:    user.ExpiresAt.Unix(),
		Role:         uint64(user.RoleID),
	}, nil
}

// hanlder to refresh token
func (s *AuthServer) RefreshToken(ctx context.Context, c *proto.RefreshTokenRequest) (*proto.AuthResponse, error) {

	// verify refresh token
	if err := service.VerifyRefreshToken(c.RefreshToken); err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid refresh token")
	}

	var user models.User
	if err := db.Db_context.Preload("Role").First(&user, "RefreshToken = ?", c.RefreshToken).Error; err != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	accessToken, err := service.GenerateAccessToken(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to generate access token")
	}

	refreshToken, err := service.GenerateRefreshToken(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to generate refresh token")
	}

	// update token in db
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	user.ExpiresIn = time.Now().Add(time.Duration(env.ACCESS_TOKEN_EXP))
	user.ExpiresAt = time.Now().Add(time.Duration(env.REFRESH_TOKEN_EXP))

	if err := db.Db_context.Save(&user).Error; err != nil {
		return nil, status.Error(codes.Internal, "Failed to save token in db")
	}

	return &proto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    user.ExpiresIn.Unix(),
		ExpiresAt:    user.ExpiresAt.Unix(),
		Role:         uint64(user.RoleID),
	}, nil
}