package service

import (
	"fmt"
	"goFinal/model"
	"goFinal/repository"

	"gorm.io/gorm"
)

type showDataService interface {
	GetAllUsers() *[]model.User
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	InsertUser(data model.User) (int64, error)
	UpdateUserByID(id int, data model.User) (int, error)
	UpdateUserByEmail(email string, data model.User) (int, error)
	// GetAllCountriesByName(name string) (*[]model.Country, error)
	// GetAllLandmarks() (*[]model.Landmark, error)
	// GetAllLandmarkByName(name string) (*[]model.Landmark, error)
	// UpdateNameAndDetail(id int, landmark model.Landmark) (int, error)
	// DeleteLandmarkById(id int) (int, error)
	// InsertLandmark(landmark *model.Landmark) (int, error)
}

func NewUsersService(gormdb *gorm.DB) showDataService {
	return showData{db: gormdb}
}

type showData struct {
	db *gorm.DB
}

func (c showData) UpdateUserByID(id int, data model.User) (int, error) {
	userRepository := repository.NewUsersRepository(c.db)
	affectedRow, err := userRepository.UpdateUserByID(id, data)
	if err != nil {
		return -1, err
	}
	return affectedRow, nil
}

func (c showData) UpdateUserByEmail(email string, data model.User) (int, error) {
	userRepository := repository.NewUsersRepository(c.db)
	affectedRow, err := userRepository.UpdateUserByEmail(email, data)
	if err != nil {
		return -1, err
	}
	return affectedRow, nil
}

func (c showData) InsertUser(data model.User) (int64, error) {
	userRepository := repository.NewUsersRepository(c.db)
	affectedRow, err := userRepository.InsertUser(data)
	if err != nil {
		return -1, err
	}
	return affectedRow, nil
}

func (c showData) GetAllUsers() *[]model.User {
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

func (c showData) GetUserById(id int) (*model.User, error) {
	userRepository := repository.NewUsersRepository(c.db)
	user, err := userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c showData) GetUserByEmail(email string) (*model.User, error) {
	userRepository := repository.NewUsersRepository(c.db)
	user, err := userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
