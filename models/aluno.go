package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"required"`
	Cpf  string `json:"cpf" validate:"len=11, number"`
	Rg   string `json:"rg" validate:"len=9, number"`
}

func ValidarDadosAluno(aluno *Aluno) error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(aluno); err != nil {
		return err
	}

	return nil
}