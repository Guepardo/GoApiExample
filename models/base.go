package models

import (
	"log"

	"github.com/gofrs/uuid"
)

type BasicModel struct {
	Token string `json:token`
}

func (model *BasicModel) GenerateToken() {
	token, err := uuid.NewV4()

	if err != nil {
		log.Fatal("Error on uuid generation")
	}

	model.Token = token.String()
}
