package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juggwow/peagolang/db"
	"github.com/juggwow/peagolang/model"
)

func ListCourses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courses, err := db.GetAllCourse()
		if err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, courses)
	}
}

func GetCourse(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}
		course, err := db.GetCourse(uint(id))
		if err != nil {
			sendError(c, http.StatusNotFound, err)
			return
		}
		c.IndentedJSON(http.StatusOK, course)
	}
}

func CreateCourse(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		trainer, err := db.GetUserByID(User(c).ID)
		if err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := trainer.Profile.IsTrainer(); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		req := new(model.Course)
		if err := c.BindJSON(req); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}
		if err := db.CreateCourse(req); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.IndentedJSON(http.StatusOK, req)
	}
}
