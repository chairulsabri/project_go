package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailabel(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, filelocation string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

//mapping struct input struct user
//Simpan struct User melalui repository'

func (s *service) IsEmailAvailabel(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

// SaveAvatar implements Service.
func (s *service) SaveAvatar(ID int, filelocation string) (User, error) {
	//dapatkan user berdasarkan ID
	//update atribut avatar file name
	//simpan perubahan file name

	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = filelocation

	updateUser, err := s.repository.Update(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No user found on with ID")
	}

	return user, nil
}
