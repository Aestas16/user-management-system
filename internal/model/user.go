package model

import (
    "errors"
)

type User struct {
    ID          uint    `gorm:"primaryKey;autoIncrement;index"`
    Username    string  `gorm:"type:varchar(80);unique;not null;" json:"username"`
    Password    string  `json:"password"`
    Email       string  `json:"email"`
}

var userNotFound = errors.New("user not found")
var userAlreadyExist = errors.New("user already exist")

func createUser(user *User) error {
    _, err := findUserByName(user.username)
    if err == nil {
        return userAlreadyExist
    }
    return db.Model(&User{}).Create(user)
}

func saveUser(user *User) error {
    return db.Save(user).Error
}

func deleteUserById(id uint) error {
    _, err := findUserById(id)
    if err != nil {
        return err
    }
    return db.Delete(&User{}, id).Error
}

func findUserByName(username string) (*User, error) {
    user := User{}
    result := db.Model(&User{}).Where("username = ?", username).First(&user)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, userNotFound
    }
    return user, result.Error
}

func findUserById(id uint) (*User, error) {
    user := User{}
    result := db.Model(&User{}).Where("id = ?", id).First(&user)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, userNotFound
    }
    return user, result.Error
}