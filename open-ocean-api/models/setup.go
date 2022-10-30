package models

import (
    "fmt"
    "github.com/gocarina/gocsv"
    _ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
    "github.com/spf13/viper"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

func SetUpModels() *gorm.DB {
    viper.AutomaticEnv()

    settingDbUser := viper.Get("POSTGRES_USER")
    settingDbPw := viper.Get("POSTGRES_PASSWORD")
    settingDbHost := viper.Get("POSTGRES_HOST")
    settingDbPort := viper.Get("POSTGRES_PORT")
    settingDb := viper.Get("POSTGRES_DB")

    postgres_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", settingDbHost, settingDbPort, settingDbUser, settingDb, settingDbPw)

    db, err := gorm.Open(postgres.Open(postgres_conname), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to DB!")
    }

    db.AutoMigrate(&User{}, &Score{}, &ScoringKey{}, &UserAnswer{})
    //Populate Table With Scoring Key Data
    firstSk := ScoringKey{}
    result := db.First(&firstSk)
    if result.RowsAffected <= 0 {
        scoringFile, err := os.OpenFile("data/scoring-key.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
        if err != nil {
            panic(err)
        }
        defer scoringFile.Close()

        sk := []*ScoringKey{}

        parseErr := gocsv.UnmarshalFile(scoringFile, &sk)
        if parseErr != nil {
            fmt.Println(err.Error())
            panic(err)
        }
        db.Create(&sk)
    }
    return db
}
