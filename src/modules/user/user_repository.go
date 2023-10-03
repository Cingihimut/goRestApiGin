package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() []User
	FindOne(id int) User
	FindByEmail(string) (*User, error)
	Save(user User) (*User, error)
	Update(user User) (*User, error)
	Delete(user User) (*User, error)
	POST(user User) (*User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (ur *UserRepositoryImpl) FindAll() []User {
	var users []User

	_ = ur.db.Find(&users)

	return users
}

func (ur *UserRepositoryImpl) FindOne(id int) User {
	var user User

	_ = ur.db.First(&user, id)

	return user
}

func (ur *UserRepositoryImpl) FindByEmail(email string) (*User, error) {
	var user User

	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepositoryImpl) Save(user User) (*User, error) {
	result := ur.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}

func (ur *UserRepositoryImpl) Update(user User) (*User, error) {
	result := ur.db.Save(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}

func (ur *UserRepositoryImpl) Delete(user User) (*User, error) {
	result := ur.db.Delete(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}

func (ur *UserRepositoryImpl) POST(user User) (*User, error) {
	result := ur.db.Save(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
