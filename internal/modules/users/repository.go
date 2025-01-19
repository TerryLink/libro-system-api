package users

import (
	"errors"
	"log"

	"gorm.io/gorm"

	"libro-system-api/internal/config"
	"libro-system-api/internal/database"
	"libro-system-api/internal/modules/util"
)

type userRepository interface {
	Login()
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (*User, error)
	AddUser(newUser User) (*User, error)
	FindByEmailOrAccountName(accountName string, email string) (*User, error)
	DeleteUser(id string) (bool, error)
	UpdateUser(updatedUser User) (*User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db database.Service) *UserRepository {
	return &UserRepository{db: db.GetDB()}
}

func (u *UserRepository) Login(accountName string, password string) (*JWTToken, error) {
	var user User
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	if err := u.db.Where("account_name = ?", user.AccountName).First(&user).Error; err != nil {
		return nil, err
	}
	if !util.ValidatePassword(user.Password, []byte(password)) {
		return nil, errors.New("invalid password")
	}
	tokenString, err := util.CreateJWT([]byte(cfg.JWTSecret), accountName)
	if err != nil {
		return nil, err
	}
	return &JWTToken{Token: tokenString}, nil
}

// Get all users
func (u *UserRepository) GetAllUsers() ([]User, error) {
	var users []User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) GetUserByID(id string) (*User, error) {
	var user User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindByEmailOrAccountName(accountName string, email string) (*User, error) {
	var user User
	if err := u.db.Where("account_name = ? or email = ?", user.AccountName, user.Email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) AddUser(newUser User) (*User, error) {
	if newUser.AccountName == "" || newUser.FirstName == "" || newUser.LastName == "" || newUser.Email == "" {
		return nil, errors.New("missing parameters")
	}
	// search duplicate AccountName or Email
	var existingUser User
	err := u.db.Where("account_name = ? or email = ?", newUser.AccountName, newUser.Email).First(&existingUser).Error
	if err == nil {
		return nil, errors.New("account name or email already exist")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err := u.db.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (u *UserRepository) DeleteUser(id string) (bool, error) {
	var existingUser User
	err := u.db.Where("id = ?", id).First(&existingUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("user not found")
	} else if err != nil {
		return false, err
	}
	// update DeletedAt
	if err := u.db.Delete(&existingUser).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserRepository) UpdateUser(updatedUser User) (*User, error) {
	var existingUser User
	if updatedUser.AccountName == "" || updatedUser.FirstName == "" || updatedUser.LastName == "" || updatedUser.Email == "" {
		return nil, errors.New("missing parameters")
	}
	err := u.db.Where("account_name = ? or email = ?", updatedUser.AccountName, updatedUser.Email).First(&existingUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	// patch user reocrd
	updateErr := u.db.Model(&existingUser).Updates(updatedUser).Error
	if updateErr != nil {
		return nil, updateErr
	}
	return &existingUser, nil
}
