package jwt

import (
	"time"

	"Auth-microservice/models"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.RegisteredClaims
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey: secretKey, tokenDuration: tokenDuration}
}

// create token and refresh token
func (m *JWTManager) GenerateToken(user *models.User) (accessToken string, refreshToken string, err error) {
	// claim information user
	claims := &UserClaims{
		ID: user.ID,
		Email: user.Email,
		Role: user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.tokenDuration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(m.secretKey))
	if err != nil {
		return "", "", err
	}
	
	refreshClaims := claims
	refreshClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(m.tokenDuration * 24))
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims)
	refreshToken, err = refreshTokenObj.SignedString([]byte(m.secretKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// verify token and return claims
func (m *JWTManager) Verify(accessToen string) (*UserClaims, error){
	token, err := jwt.ParseWithClaims(accessToen, &UserClaims{}, func(token *jwt.Token) (interface{}, error){
		return []byte(m.secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claim, ok := token.Claims.(*UserClaims)

	if !ok {
		return nil, err
	}

	return claim, nil
}


