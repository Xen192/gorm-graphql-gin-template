package auth

import (
	"mygpt/models"
	"mygpt/pkg/infrastructure/datastore"

	"github.com/oklog/ulid/v2"
)

func SeedUser(clrk *models.ClerkUser) error {
	db := datastore.GetInstance()
	tx := db.Begin()
	u, err := datastore.ClerkClient.Users().Read(clrk.ID)
	if err != nil {
		return err
	}

	clerkUser := models.ClerkUser{
		ID: u.ID,
	}
	if len(u.EmailAddresses) > 0 {
		clerkUser.LinkedIdentity = u.EmailAddresses[0].EmailAddress
		clrk.LinkedIdentity = clerkUser.LinkedIdentity
	}
	res := tx.Create(clerkUser)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	user := models.User{
		ID:              ulid.Make().String(),
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Status:          models.StatusActive,
		ProfileImageURL: &u.ProfileImageURL,
	}
	if len(u.EmailAddresses) > 0 {
		user.Email = &u.EmailAddresses[0].EmailAddress
	}
	res = tx.Save(&user)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	tx.Commit()
	return nil
}
