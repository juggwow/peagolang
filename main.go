package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/juggwow/peagolang/db"
	"github.com/juggwow/peagolang/handler"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/studentprofile", handler.RequireUser(db), handler.AddStudentProfile(db)) //
	r.POST("/trainerprofile", handler.RequireUser(db), handler.AddTrainerProfile(db)) //
	r.GET("/profile", handler.RequireUser(db), handler.GetProfile(db))                //

	r.POST("/courses", handler.RequireUser(db), handler.CreateCourse(db))
	r.GET("/courses", handler.ListCourses(db))   //
	r.GET("/courses/:id", handler.GetCourse(db)) //

	r.POST("/classes", handler.RequireUser(db), handler.CreateClasses(db))       //
	r.GET("/classes/:id", handler.ListClassesByCourseID(db))                     //
	r.DELETE("/classes/:id", handler.RequireUser(db), handler.DeleteClasses(db)) //
	r.PATCH("/classes", handler.RequireUser(db), handler.EditClasses(db))

	r.POST("/enrollments", handler.RequireUser(db), handler.EnrollClass(db)) //

	r.POST("/register", handler.Register(db)) //
	r.POST("/login", handler.Login(db))       //

	r.GET("/test", Gettest()) //

	r.Run(":8624")
}

func Error(c *gin.Context, status int, err error) {
	log.Println(err)
	c.JSON(status, gin.H{
		"message": err.Error(),
	})
}

func Gettest() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.IndentedJSON(http.StatusOK, gin.H{
			"say": "Hello Hello Hello 55555",
		})

	}
}
