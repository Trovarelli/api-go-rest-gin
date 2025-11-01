package models

import (
    "time"
    "gorm.io/gorm"
)

type Aluno struct {
    Id        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
    Nome      string         `json:"nome"`
    CPF       string         `json:"cpf"`
    RG        string         `json:"rg"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
