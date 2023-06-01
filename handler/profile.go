package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juggwow/peagolang/db"
	"github.com/juggwow/peagolang/model"
)

func AddStudentProfile(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(model.Profile)

		_, err := db.GetStudent(User(c).ID)
		if err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := c.BindJSON(req); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		req.ID = User(c).ID

		if err := req.IsStudent(); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := db.SaveProfile(req); err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func AddTrainerProfile(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(model.Profile)

		_, err := db.GetStudent(User(c).ID)
		if err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := c.BindJSON(req); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		req.ID = User(c).ID

		if err := req.IsTrainer(); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := db.SaveProfile(req); err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func GetProfile(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(User(c).ID)
		userData, err := db.GetUserByID(User(c).ID)
		if err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		profile := userData.Profile
		c.IndentedJSON(http.StatusOK, profile)
	}
}
