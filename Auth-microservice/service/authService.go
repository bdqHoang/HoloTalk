package service

import (
	"Auth-microservice/config"
	"Auth-microservice/dto"
	"Auth-microservice/models"
	"errors"
	"net/mail"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var env = config.LoadEnv()

// handler to hash password
func HashPassord(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// handler to verify password
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// handler to generate token
func GenerateAccessToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"role":     user.Role.Name,
		"user_id":  user.ID,
		"exp":      time.Now().Add(env.ACCESS_TOKEN_EXP).Unix(),
	})
	return token.SignedString(env.JWT_SECRET)
}

// handler to generate refresh token
func GenerateRefreshToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"role":     user.Role.Name,
		"user_id":  user.ID,
		"exp":      time.Now().Add(env.REFRESH_TOKEN_EXP).Unix(),
	})
	return token.SignedString(env.JWT_SECRET)
}

// handel to verify token
func VerifyRefreshToken(token string)  error {
	claims := &jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.JWT_SECRET), nil
	})

	if err != nil || (*claims)["user_id"] == nil {
		return errors.New("Invalid refresh token")
	}

	return nil
}


// validate register request
func ValidateRegisterRequest(userReq dto.RegisterRequest) error {

	// validate email
	if userReq.Email == ""{
		return errors.New("Email is required")
	}

	_,err := mail.ParseAddress(userReq.Email)
	if err != nil {
		return errors.New("Invalid email")
	}

	// validate phone
	if userReq.Phone == ""{
		return errors.New("Phone is required")
	}
	
	e164Regex := `^\+[1-9]\d{1,14}$`
	re := regexp.MustCompile(e164Regex)
	if !re.MatchString(userReq.Phone) {
		return errors.New("Invalid phone number")
	}

	// validate name
	if userReq.Name == ""{
		return errors.New("Name is required")
	}

	// validate date of birth
	if userReq.DateOfBirth == ""{
		return errors.New("Date of birth is required")
	}

	// check invalid date of birth
	_, err = time.Parse("02-01-2006", userReq.DateOfBirth)
	if err != nil {
		return errors.New("Invalid date of birth")
	}

	// validate password
	if userReq.Password == ""{
		return errors.New("Password is required")
	}

	// validate confirm password
	if userReq.ConfirmPassword == ""{
		return errors.New("Confirm password is required")
	}

	if userReq.Password != userReq.ConfirmPassword {
		return errors.New("Password and confirm password do not match")
	}

	return nil
}

// validate login request
func ValidateLoginRequest(userReq dto.LoginRequest) error {

	// validate username
	if userReq.Username == ""{
		return errors.New("Username is required")
	}	

	// validate password
	if userReq.Password == ""{
		return errors.New("Password is required")
	}

	return nil
}