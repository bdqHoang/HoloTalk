package app

import (
	"Auth-microservice/Interface/IRepository"
	"Auth-microservice/Interface/IService"
	"Auth-microservice/jwtToken"
	"Auth-microservice/models"
	"errors"
	"time"

	"github.com/bdqHoang/HoloTalk/shared/dto"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo IRepository.IUserRepo
	jwtManager *jwt.JWTManager
}


func NewAuthService(userRepo IRepository.IUserRepo, jwtManager *jwt.JWTManager) IService.IAuthService {
	return &authService{userRepo: userRepo , jwtManager: jwtManager}
}

func (s *authService) Register(userReq dto.RegisterRequest) (*models.User, error){
	_, err := s.userRepo.GetByEmail(userReq.Email)

	// check existed email
	if err == nil {
		return nil, errors.New("Email already exists")
	}

	// hash passwork password+email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password+userReq.Email), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name: userReq.Name,
		Email: userReq.Email,
		Phone: userReq.Phone,
		DateOfBirth: userReq.DateOfBirth,
		PasswordHash: string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.userRepo.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(loginReq dto.LoginRequest) (string, string, error){
	user, err := s.userRepo.GetByEmail(loginReq.Email)
	if err != nil{
		return "","", errors.New("Email or password incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password+ loginReq.Email))
	if err != nil {
		return "","", errors.New("Email or password incorrect")
	}

	accessToken, refreshToken, err := s.jwtManager.GenerateToken(user)
	if err != nil{
		return "","", err
	}

	return accessToken,refreshToken, nil
}