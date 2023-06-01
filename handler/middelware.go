package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juggwow/peagolang/db"
)

func RequireUser(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get header
		header := c.GetHeader("Authorization")
		header = strings.TrimSpace(header)

		//if not found > return err http.status 401
		min := len("Bearer ")
		if len(header) <= min {
			sendError(c, http.StatusUnauthorized, errors.New("token is require"))
			return
		}

		token := header[min:]
		claims, err := verifyToken(token)
		if err != nil {
			sendError(c, http.StatusUnauthorized, err)
			return

		}

		user, err := db.GetUserByID(claims.UserID)
		if err != nil {
			sendError(c, http.StatusUnauthorized, err)
			return
		}

		SetUser(c, user)

		// Verify token
		// IF not valid -> err http 401
		// Find user from token by ID
		// If not found -> 401
		// setuser
		// else valid next() to function
	}
}
