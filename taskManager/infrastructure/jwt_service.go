package infrastructure

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(userId, username, role string) (string, error)
}

type jwtService struct {
	secretKey string
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET_KEY")
}

func NewJwtService() JwtService {
	secret := GetJWTSecret()
	return &jwtService{secretKey: secret}
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func (jwtService *jwtService) GenerateToken(userId, username, role string) (string, error) {
	claims := Claims{
		UserID:   userId,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if jwtService.secretKey == "" {
		jwtService.secretKey = os.Getenv("JWT_SECRET")
	}
	return token.SignedString([]byte(jwtService.secretKey))
}
