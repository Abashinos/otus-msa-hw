package schemas

import (
	"github.com/Abashinos/otus-msa-hw/app/pkg/models"
)

type UserSchemaBase struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserSchemaUpdate UserSchemaBase

func (u *UserSchemaUpdate) ToModel() *models.User {
	return &models.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

type UserSchemaCreate UserSchemaBase

func (u *UserSchemaCreate) ToModel() *models.User {
	return &models.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

type UserSchemaGet struct {
	UserSchemaBase
	ID uint `json:"id"`
}

func (u UserSchemaGet) FromModel(userModel *models.User) UserSchemaGet {
	u.ID = userModel.ID
	u.FirstName = userModel.FirstName
	u.LastName = userModel.LastName
	return u
}
