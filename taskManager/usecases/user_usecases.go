package usecases

import (
	"errors"
	"taskManager/domain"
	"taskManager/infrastructure"
	"taskManager/repositories"
	"taskManager/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCases interface {
	Register(input domain.UserInput) (*domain.User, error)
	LoginUser(input domain.UserInput) (string, error)
	PromoteUser(id primitive.ObjectID) error
}

type userUseCase struct {
	userRepository  repositories.UserRepository
	jwtService      infrastructure.JwtService
	passwordService infrastructure.PasswordService
}

func NewUserUseCase(jwtService infrastructure.JwtService, userRepository repositories.UserRepository, passwordService infrastructure.PasswordService) UserUseCases {
	return &userUseCase{userRepository: userRepository, jwtService: jwtService, passwordService: passwordService}
}

func (userUseCase *userUseCase) Register(input domain.UserInput) (*domain.User, error) {
	existingUser, _ := userUseCase.userRepository.FindUserByUsername(input.Username)

	if existingUser.Username != "" {
		return nil, errors.New("user with username already exists")
	}

	encodedPassword, err := userUseCase.passwordService.EncodePassword(input.Password)

	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username: input.Username,
		Password: encodedPassword,
		Role:     types.User,
	}

	count, err := userUseCase.userRepository.Count()

	if count == 0 {
		user.Role = types.Admin
	}

	if err := userUseCase.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (userUseCase *userUseCase) LoginUser(input domain.UserInput) (string, error) {
	existingUser, err := userUseCase.userRepository.FindUserByUsername(input.Username)
	if err != nil || existingUser.Username == "" {
		return "", errors.New("invalid credentials")
	}

	err = userUseCase.passwordService.ComparePassword(input.Password, existingUser.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	roleString := string(existingUser.Role)

	token, err := userUseCase.jwtService.GenerateToken(existingUser.ID.Hex(), existingUser.Username, roleString)

	return token, nil
}

func (userUseCase *userUseCase) PromoteUser(id primitive.ObjectID) error {
	err := userUseCase.userRepository.UpdateById(id)
	return err
}
