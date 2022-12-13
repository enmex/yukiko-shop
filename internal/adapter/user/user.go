package user

import (
	spec "yukiko-shop/internal/generated/spec/auth"
	"yukiko-shop/internal/repository/ent"
)

func ConvertEntToDomain(userEnt *ent.User) *spec.User {
	return &spec.User{
		Id:       userEnt.ID.String(),
		Email:    userEnt.Email,
		FirstName: userEnt.FirstName,
		LastName: userEnt.LastName,
		Password: userEnt.Password,
	}
}