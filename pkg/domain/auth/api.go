package auth

import (
	"context"
	"mygpt/model"
	"mygpt/model_struct"
	"mygpt/pkg/infrastructure/datastore"
	"mygpt/pkg/utils"
	"mygpt/query"

	"github.com/oklog/ulid/v2"
)

func SeedUser(clrk *model.ClerkUser) error {
	tx := query.Q.Begin()
	u, err := datastore.ClerkClient.Users().Read(clrk.ID)
	if err != nil {
		return err
	}

	clerkUser := model.ClerkUser{
		ID: u.ID,
	}
	if len(u.EmailAddresses) > 0 {
		clerkUser.LinkedIdentity = u.EmailAddresses[0].EmailAddress
		clrk.LinkedIdentity = clerkUser.LinkedIdentity
	}
	res := tx.ClerkUser.WithContext(context.Background()).Create(&clerkUser)
	if res != nil {
		tx.Rollback()
		return res
	}

	user := model.User{
		ID:              ulid.Make().String(),
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Status:          utils.Pointer(model_struct.UserStatusActive),
		ProfileImageURL: &u.ProfileImageURL,
	}
	if len(u.EmailAddresses) > 0 {
		user.Email = &u.EmailAddresses[0].EmailAddress
	}
	res = tx.User.WithContext(context.Background()).Create(&user)
	if res != nil {
		tx.Rollback()
		return res
	}

	tx.Commit()
	return nil
}
