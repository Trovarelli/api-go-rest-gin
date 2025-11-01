package database

import (
    "os"

    "api-go-rest-gin/src/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func DatabaseConnect() (*gorm.DB, error) {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        dsn = "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
    }
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    if err := db.AutoMigrate(&models.Aluno{}); err != nil {
        return nil, err
    }

    return db, nil
}
