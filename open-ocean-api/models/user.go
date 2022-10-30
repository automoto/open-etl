package models

import (
    "database/sql"
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    TestID uint `json:"test_user_id"`
    Sex int `json:"sex"`
    Age string `json:"age"`
    Country string `json:"country"`
    Scored bool `json:"scored" gorm:"default:False"`
    Email sql.NullString `json:"email "`
    UserAnswers []UserAnswer
    Scores []Score
}

type UserCRUD struct {
    db *gorm.DB
}

func (uc *UserCRUD) CreateUser(input *User) (err error) {
    uErr := uc.db.Create(&input).Error
    if uErr != nil {
        return uErr
    }
    return nil
}

func (uc *UserCRUD) UpdateUser(id int, newUserData *User) (err error){
    user := User{}
    uErr := uc.db.Where("id = ?", id).First(&user).Error
    if uErr != nil {
        return uErr
    }
    uErr = uc.db.Model(&user).Updates(newUserData).Error
    if uErr != nil {
        return uErr
    }
    return nil
}

func (uc *UserCRUD) DeleteUser(id int) (err error){
    user := User{}
    uErr := uc.db.Where("id = ?", id).First(&user).Error
    if uErr != nil {
        return uErr
    }
    uErr = uc.db.Delete(user).Error
    if uErr != nil {
        return uErr
    }
    return nil
}

func (uc *UserCRUD) GetUser(id int) (u User, err error){
    user := User{}
    uErr := uc.db.Where("id = ?", id).First(&user).Error
    if uErr != nil {
        return user, uErr
    }
    return user, nil
}
