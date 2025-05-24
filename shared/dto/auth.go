package dto

// requset struct
type RegisterRequest struct {
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Name            string `json:"name"`
	DateOfBirth     string `json:"dateOfBirth"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Role            uint   `json:"role"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

//response struct
type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int64  `json:"expiresIn"`
	ExpiresAt    int64  `json:"expiresAt"`
	Role         uint   `json:"role"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}