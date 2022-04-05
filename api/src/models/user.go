package models

import (
	"api/src/services/passwordServices"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

//User representa um usuario na tabela
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}
	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Insira todos os campos! (Nome)")
	}
	if user.Nick == "" {
		return errors.New("Insira todos os campos! (Nick)")
	}
	if user.Email == "" {
		return errors.New("Insira todos os campos! (Email)")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O e-mail é inválido")
	}

	if step == "register" && user.Password == "" {
		if user.Password == "" {
			return errors.New("Insira todos os campos! (Senha)")
		}
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashPass, err := passwordServices.HashPassword(user.Password)

		if err != nil {
			return err
		}
		user.Password = string(hashPass)
	}
	return nil
}
