package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Create(user *User) (*User, error) {
	if err := u.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Get(userID uint) (user *User, err error) {
	err = u.DB.Where(userID).First(&user).Error
	return
}

func (u *UserRepository) Update(userId uint, user *User) (*User, error) {
	if _, err := u.Get(userId); err != nil {
		return nil, err
	}

	user.ID = userId
	if err := u.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return u.Get(userId)
}

func (u *UserRepository) Delete(userID uint) (err error) {
	err = u.DB.Delete(&User{}, userID).Error
	return
}
