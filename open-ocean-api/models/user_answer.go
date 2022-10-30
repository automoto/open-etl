package models

import "github.com/jinzhu/gorm"

type UserAnswer struct{
    gorm.Model
    QuestionNum int `json:"question_num"`
    Answer int `json:"answer"`
    TestID int
    UserId uint
}