package db

import (
	"log"
	"os"
	"time"

	"github.com/juggwow/peagolang/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	db *gorm.DB
}

func NewDB() (*DB, error) {
	url := os.Getenv("DATABASE_URL")
	//url := os.Getenv("DATABASE_URL")
	//url := "host=postgresql user=peagolang password=supersecret dbname=peagolang port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	if err := db.Migrator().AutoMigrate(
		&User{},
		&Course{},
		&Class{},
		&ClassStudent{},
		&Profile{},
	); err != nil {
		log.Fatal(err)
	}

	return &DB{db: db}, nil
}

type Course struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
}

type Class struct {
	ID        uint `gorm:"primaryKey"`
	CourseID  uint
	Course    Course
	TrainerID uint
	Trainer   User
	Start     time.Time
	End       time.Time
	Seats     int
	Students  []ClassStudent
}

type ClassStudent struct {
	ID        uint `gorm:"primaryKey"`
	ClassID   uint
	StudentID uint
	Student   User
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Profile  Profile
}

type Profile struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Firstname string
	Lastname  string
	Role      string
	Company   string
	Mobileno  string
}

func (db *DB) Reset() error {
	err := db.db.Migrator().AutoMigrate(
		&User{},
		&Course{},
		&Class{},
		&ClassStudent{},
		&Profile{},
	)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (db *DB) AutoMigrate() error {
	err := db.db.Migrator().DropTable(
		&User{},
		&Course{},
		&Class{},
		&ClassStudent{},
		&Profile{},
	)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (db *DB) CreateUser(u *model.User) error {
	user := User{
		Username: u.Username,
		Password: u.Password,
	}
	if err := db.db.Create(&user).Error; err != nil {
		return err
	}
	u.ID = user.ID
	return nil
}

func (db *DB) CreateCourse(c *model.Course) error {
	course := Course{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
	if err := db.db.Save(&course).Error; err != nil {
		return err
	}
	c.ID = course.ID
	return nil
}

func (db *DB) GetCourse(id uint) (*model.Course, error) {
	var course Course
	if err := db.db.First(&course, id).Error; err != nil {
		return nil, err
	}
	return &model.Course{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
	}, nil
}

func (db *DB) GetAllCourse() ([]model.Course, error) {
	var courses []Course
	if err := db.db.Find(&courses).Error; err != nil {
		return nil, err
	}

	result := []model.Course{}
	for _, course := range courses {
		result = append(result, model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
		})
	}

	return result, nil
}

func (db *DB) SaveClass(cls *model.Class) error {
	class := Class{
		CourseID:  cls.Course.ID,
		TrainerID: cls.Trainer.ID,
		Start:     cls.Start,
		End:       cls.End,
		Seats:     cls.Seats,
	}
	if err := db.db.Save(&class).Error; err != nil {
		return err
	}
	cls.ID = class.ID
	return nil
}

func (db *DB) DeleteClass(cls *model.Class) error {
	class := Class{
		ID:        cls.ID,
		CourseID:  cls.Course.ID,
		TrainerID: cls.Trainer.ID,
		Start:     cls.Start,
		End:       cls.End,
		Seats:     cls.Seats,
	}
	if err := db.db.Delete(&class).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) EditClass(cls *model.Class) error {
	class := Class{
		ID: cls.ID,
	}
	//db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	if err := db.db.Model(&class).Updates(map[string]interface{}{"start": cls.Start, "end": cls.End, "seats": cls.Seats}).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) GetAllClassbyCourseID(id uint) ([]model.Class, error) {
	var class []Class
	err := db.db.Preload("Course").Preload("Trainer").Preload("Students.Student").Where("course_id = ?", id).Find(&class).Error
	if err != nil {
		return nil, err
	}
	result := []model.Class{}
	for _, cl := range class {
		students := []model.Student{}
		for _, stu := range cl.Students {
			students = append(students, model.Student{
				ID:   stu.StudentID,
				Name: stu.Student.Username,
			})
		}
		result = append(result, model.Class{
			ID: cl.ID,
			Course: model.Course{
				ID:          cl.Course.ID,
				Name:        cl.Course.Name,
				Description: cl.Course.Description,
			},
			Trainer: model.Trainer{
				ID:   cl.Trainer.ID,
				Name: cl.Trainer.Username,
			},
			Start:    cl.Start,
			End:      cl.End,
			Seats:    cl.Seats,
			Students: students,
		})
	}
	//ส่งออก เป็น Model Class
	// type Class struct {
	// 	ID       uint
	// 	Course   Course
	// 	Trainer  Trainer
	// 	Start    time.Time
	// 	End      time.Time
	// 	Seats    int
	// 	Students []Student
	// }
	return result, nil
}

func (db *DB) GetClass(id uint) (*model.Class, error) {
	//ดึงข้อมูล Class จาก id
	var class Class
	err := db.db.Preload("Course").
		Preload("Trainer").
		Preload("Students.Student"). //ดึงข้อมูล student ทั้งหมด
		First(&class, id).Error
	if err != nil {
		return nil, err
	}

	//จัดเก็บ student เป็น Model Student
	// type Student struct {
	// 	ID   uint
	// 	Name string
	// }
	students := []model.Student{}
	for _, stu := range class.Students {
		students = append(students, model.Student{
			ID:   stu.StudentID,
			Name: stu.Student.Username,
		})
	}

	//ส่งออก เป็น Model Class
	// type Class struct {
	// 	ID       uint
	// 	Course   Course
	// 	Trainer  Trainer
	// 	Start    time.Time
	// 	End      time.Time
	// 	Seats    int
	// 	Students []Student
	// }
	return &model.Class{
		ID: class.ID,
		Course: model.Course{
			ID:          class.Course.ID,
			Name:        class.Course.Name,
			Description: class.Course.Description,
		},
		Trainer: model.Trainer{
			ID:   class.Trainer.ID,
			Name: class.Trainer.Username,
		},
		Start:    class.Start,
		End:      class.End,
		Seats:    class.Seats,
		Students: students,
	}, nil
}

func (db *DB) GetStudent(id uint) (*model.Student, error) {
	var student User
	if err := db.db.Preload("Profile").First(&student, id).Error; err != nil {
		return nil, err
	}
	return &model.Student{
		ID:   student.ID,
		Name: student.Username,
		Profile: model.Profile{
			Firstname: student.Profile.Firstname,
			Lastname:  student.Profile.Lastname,
			Role:      student.Profile.Role,
			Company:   student.Profile.Company,
			Mobileno:  student.Profile.Mobileno,
		},
	}, nil
}

func (db *DB) CreateClassStudent(studentID uint, classID uint) error {
	classStudent := ClassStudent{
		StudentID: studentID,
		ClassID:   classID,
	}
	return db.db.Create(&classStudent).Error
}

func (db *DB) DeleteClassStudentbyClassID(classID uint) error {
	return db.db.Where("class_id = ?", classID).Delete(&ClassStudent{}).Error
}

func (db *DB) GetUser(username string) (*model.User, error) {
	var user User

	if err := db.db.Preload("Profile").First(&user, "Username=?", username).Error; err != nil {
		return nil, err
	}

	//ส่งออก User เป็น Model User
	// type User struct {
	// 	ID       uint
	// 	Username string
	// 	Password string
	// }
	return &model.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Profile: model.Profile{
			Firstname: user.Profile.Firstname,
			Lastname:  user.Profile.Lastname,
			Role:      user.Profile.Role,
			Company:   user.Profile.Company,
			Mobileno:  user.Profile.Mobileno,
		},
	}, nil
}

func (db *DB) GetUserByID(id uint) (*model.User, error) {
	var user User

	if err := db.db.Preload("Profile").First(&user, id).Error; err != nil {
		return nil, err
	}

	//ส่งออก User เป็น Model User
	// type User struct {
	// 	ID       uint
	// 	Username string
	// 	Password string
	// }
	return &model.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Profile: model.Profile{
			Firstname: user.Profile.Firstname,
			Lastname:  user.Profile.Lastname,
			Role:      user.Profile.Role,
			Company:   user.Profile.Company,
			Mobileno:  user.Profile.Mobileno,
		},
	}, nil
}

func (db *DB) SaveProfile(profile *model.Profile) error {
	var user User
	if err := db.db.Preload("Profile").First(&user, profile.ID).Error; err != nil {
		return err
	}

	userProfile := Profile{
		UserID:    profile.ID,
		Firstname: profile.Firstname,
		Lastname:  profile.Lastname,
		Role:      profile.Role,
		Company:   profile.Company,
		Mobileno:  profile.Mobileno,
	}

	if user.Profile.ID == 0 {
		return db.db.Save(&userProfile).Error
	} else {
		return db.db.Model(&Profile{}).Where("id=?", user.Profile.ID).Updates(&userProfile).Error
	}

}
