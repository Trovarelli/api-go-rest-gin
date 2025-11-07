package models

import (
	"time"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	Id        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Nome      string         `json:"nome" validate:"nonzero"`
	CPF       string         `json:"cpf" validate:"len=9, regexp=^[0-9]*$"`
	RG        string         `json:"rg" validate:"len=11, regexp=^[0-9]*$"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func AlunoValidator(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}

	return nil
}
