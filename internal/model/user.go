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

func createUser(user *User) error {
    _, err := findByName(user.username)
    if err == nil {
        return errors("user already exist")
    }
    return db.Model(&User{}).Create(user);
}

func saveUser(user *User) error {
    return db.Save(user).Error;
}

func findByName(username string) (*User, error) {
    user := User{}
    result := db.Model(&User{}).Where("username = ?", username).First(&user)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, errors("user not found")
    }
    return user, result.Error
}

func findById(id uint) (*User, error) {
    user := User{}
    result := db.Model(&User{}).Where("id = ?", id).First(&user)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, errors("user not found")
    }
    return user, result.Error
}