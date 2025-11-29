package infrastructure

import "golang.org/x/crypto/bcrypt"

type PasswordService interface {
	EncodePassword(password string) (string, error)
	ComparePassword(password, encodedPassword string) error
}

type bcryptService struct{}

func NewPasswordService() PasswordService {
	return &bcryptService{}
}

func (s *bcryptService) EncodePassword(password string) (string, error) {
	encodedPassword, passwordErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encodedPassword), passwordErr
}

func (s *bcryptService) ComparePassword(password, encodedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(password))
}
