package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input ResgisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) RegisterUser(input ResgisterUserInput) (User, error) {
	user := User{}

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	user.Role = "user"

	if err != nil {
		return user, err
	}
	user.PasswordHash = string(PasswordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}