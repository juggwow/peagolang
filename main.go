package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/juggwow/peagolang/db"
	"github.com/juggwow/peagolang/handler"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/courses", handler.ListCourses(db))
	r.GET("/courses/:id", handler.GetCourse(db))
	r.POST("/courses", handler.CreateCourse(db))
	r.POST("/classes", handler.CreateClasses(db))
	r.POST("/enrollments", handler.EnrollClass(db))
	r.POST("/register", handler.Register(db))
	r.POST("/login", handler.Login(db))

	r.Run(":8080")
}

func Error(c *gin.Context, status int, err error) {
	log.Println(err)
	c.JSON(status, gin.H{
		"message": err.Error(),
	})
}
