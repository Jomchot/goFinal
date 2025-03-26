package service

import (
	"fmt"
	"goFinal/model"
	"goFinal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type showDataService interface {
	GetAllUsers() *[]model.Customer
	GetUserById(id int) (*model.Customer, error)
	PostAuthLogin(email string, password string) (*DataAuth, error)
	InsertUser(data model.Customer) (int64, error)
	UpdateUserByID(id int, data model.Customer) (int, error)
	UpdateUserByEmail(email string, data model.Customer) (int, error)
}

func NewUsersService(gormdb *gorm.DB) showDataService {
	return showData{db: gormdb}
}

type showData struct {
	db *gorm.DB
}

type DataAuth struct {
	CustomerID  uint
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (c showData) UpdateUserByID(id int, data model.Customer) (int, error) {
	userRepository := repository.NewUsersRepository(c.db)
	affectedRow, err := userRepository.UpdateUserByID(id, data)
	if err != nil {
		return -1, err
	}
	return affectedRow, nil
}

func (c showData) UpdateUserByEmail(email string, data model.Customer) (int, error) {
	userRepository := repository.NewUsersRepository(c.db)
	affectedRow, err := userRepository.UpdateUserByEmail(email, data)
	if err != nil {
		return -1, err
	}
	return affectedRow, nil
}
func (c showData) PostAuthLogin(email string, password string) (*DataAuth, error) {
	userRepository := repository.NewUsersRepository(c.db)
	user, err := userRepository.GetUserByEmail(email)
	fmt.Printf(user.Password)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		err = fmt.Errorf("password incorrect")
		return nil, err
	}

	userResponse := &DataAuth{
		CustomerID:  user.CustomerID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
	return userResponse, nil
}
func (c showData) InsertUser(data model.Customer) (int64, error) {
	userRepository := repository.NewUsersRepository(c.db)
	affectedRow, err := userRepository.InsertUser(data)
	if err != nil {
		return -1, err
	}
	return affectedRow, nil
}

func (c showData) GetAllUsers() *[]model.Customer {
	userRepository := repository.NewUsersRepository(c.db)
	dataUser, err := userRepository.GetAllUsers()

	if err != nil {
		panic(err)
	}

	for _, v := range *dataUser {
		fmt.Printf("%v", v)
	}
	return dataUser
}

func (c showData) GetUserById(id int) (*model.Customer, error) {
	userRepository := repository.NewUsersRepository(c.db)
	user, err := userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
