package middleware

import (
	"context"
	"mygpt/pkg/adapter/ctxparser"
	"mygpt/pkg/domain/auth"
	"mygpt/pkg/infrastructure/datastore"
	"mygpt/query"
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

		u, err := query.ClerkUser.WithContext(c.Request.Context()).Where(query.ClerkUser.ID.Eq(claims.Subject)).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				err := auth.SeedUser(u)
				if err != nil {
					logrus.Error(err)
					c.AbortWithStatusJSON(http.StatusInternalServerError, "Failed Creating User")
					return
				}
			} else {
				logrus.Error(err)
			}
		}

		user, err := query.User.WithContext(c.Request.Context()).Where(query.User.Email.Eq(u.LinkedIdentity)).First()
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Failed Getting User")
			return
		}

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), ctxparser.CTXUser, user))
		c.Next()
	}
}
