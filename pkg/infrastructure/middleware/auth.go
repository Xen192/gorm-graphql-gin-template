package middleware

import (
	"context"
	"mygpt/models"
	"mygpt/pkg/adapter/ctxparser"
	"mygpt/pkg/domain/auth"
	"mygpt/pkg/infrastructure/datastore"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := datastore.ClerkClient.VerifyToken(strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer "))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		db := datastore.GetInstance()
		u := models.ClerkUser{ID: claims.Subject}
		res := db.First(&u)
		if res.Error != nil {
			if res.Error == gorm.ErrRecordNotFound {
				err := auth.SeedUser(&u)
				if err != nil {
					logrus.Error(err)
					c.AbortWithStatusJSON(http.StatusInternalServerError, "Failed Creating User")
					return
				}
			} else {
				logrus.Error(err)
			}
		}

		user := models.User{Email: &u.LinkedIdentity}
		res = db.Where(&user).Find(&user)
		if res.Error != nil {
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Failed Getting User")
			return
		}

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), ctxparser.CTXUser, &user))
		c.Next()
	}
}
