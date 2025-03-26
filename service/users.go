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
	PostAuthLogin(email string, password string) (*DataAuth, error)
	InsertUser(data model.Customer) (int64, error)
	UpdatePasswordByEmail(email string, password string) (int, error)
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

func (c showData) UpdatePasswordByEmail(email string, password string) (int, error) {
	userRepository := repository.NewUsersRepository(c.db)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	// แปลง []byte -> string
	hashedPassword := string(hashedPasswordBytes)
	//update
	affectedRow, err := userRepository.UpdatePasswordByEmail(email, hashedPassword)

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
