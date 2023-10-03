package user

import (
	"strconv"

	dto "github.com/Cingihimut/goRestApiGin.git/src/modules/user/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	GetAll() []User
	GetById(id int) User
	Create(ctx *gin.Context) (*User, error)
	Update(ctx *gin.Context) (*User, error)
	Delete(ctx *gin.Context) (*User, error)
	Login(email, password string) (*User, error)
	Register(string, string, string) (*User, error)
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (us *UserServiceImpl) GetAll() []User {
	return us.userRepository.FindAll()
}

func (us *UserServiceImpl) GetById(id int) User {
	return us.userRepository.FindOne(id)
}

func (us *UserServiceImpl) Create(ctx *gin.Context) (*User, error) {
	var input dto.CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	user := User{
		Name:  input.Name,
		Email: input.Email,
	}

	result, err := us.userRepository.Save(user)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (us *UserServiceImpl) Update(ctx *gin.Context) (*User, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.UpdateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	user := User{
		ID:    int64(id),
		Name:  input.Name,
		Email: input.Email,
	}

	result, err := us.userRepository.Update(user)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (us *UserServiceImpl) Delete(ctx *gin.Context) (*User, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user := User{
		ID: int64(id),
	}

	result, err := us.userRepository.Delete(user)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (us *UserServiceImpl) Login(email, password string) (*User, error) {
	var input dto.LoginInput
	input.Email = email
	input.Password = password

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	newUser := User{
		Name:  input.Name,
		Email: input.Email,
	}

	result, err := us.userRepository.POST(newUser)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *UserServiceImpl) Register(email, name, password string) (*User, error) {
	var input dto.CreateUserInput

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	user := User{
		Name:  input.Name,
		Email: input.Email,
	}

	result, err := us.userRepository.Save(user)

	if err != nil {
		return nil, err
	}

	return result, nil

}
