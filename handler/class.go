package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juggwow/peagolang/db"
)

type ClassReq struct {
	CourseID uint      `json:"course_id"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Seats    int       `json:"seats"`
}

type EditClassReq struct {
	ClassID uint      `json:"class_id"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Seats   int       `json:"seats"`
}

func EditClasses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(EditClassReq)
		if err := c.BindJSON(req); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		class, err := db.GetClass(req.ClassID)
		if err := class.EditClass(req.Start, req.End, req.Seats); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		trainer, err := db.GetUserByID(User(c).ID)
		if err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := trainer.Profile.IsTrainer(); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := class.DeleteByTrainer(trainer.ID); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := db.EditClass(class); err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)

	}
}

func CreateClasses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(ClassReq)
		if err := c.BindJSON(req); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}
		course, err := db.GetCourse(req.CourseID)
		if err != nil {
			sendError(c, http.StatusNotFound, err)
			return
		}
		class, err := course.CreateClass(req.Start, req.End)
		if err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}
		if err := class.SetSeats(req.Seats); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		trainer, err := db.GetUserByID(User(c).ID)
		if err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := trainer.Profile.IsTrainer(); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		class.Trainer.ID = trainer.ID

		if err := db.SaveClass(class); err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func DeleteClasses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		class, err := db.GetClass(uint(id))
		if err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		trainer, err := db.GetUserByID(User(c).ID)
		if err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := trainer.Profile.IsTrainer(); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := class.DeleteByTrainer(trainer.ID); err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}

		if err := db.DeleteClassStudentbyClassID(class.ID); err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}

		if err := db.DeleteClass(class); err != nil {
			sendError(c, http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func ListClassesByCourseID(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			sendError(c, http.StatusBadRequest, err)
			return
		}
		class, err := db.GetAllClassbyCourseID(uint(id))
		if err != nil {
			sendError(c, http.StatusNotFound, err)
			return
		}
		c.IndentedJSON(http.StatusOK, class)

	}
}
