package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price float64
	Stock int
}

type Order struct {
	ID       uint `gorm:"primaryKey"`
	Products []ProductOrder
}

type ProductOrder struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	Product   Product
	OrderID   uint
	Order     Order
	Amount    int
}

type User struct {
	gorm.Model
	Username string
	Profile  StudentProfile
}

type StudentProfile struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	CompanyName string
	JobTitle    string
	Level       string
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

func main() {
	url := "host=localhost user=peagolang password=supersecret dbname=peagolang port=54329 sslmode=disable"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().DropTable(
		&Product{},
		&Order{},
		&User{},
		&StudentProfile{},
		&ProductOrder{},
		&Course{},
		&Class{},
		&ClassStudent{},
	)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrator().AutoMigrate(
		&Product{},
		&Order{},
		&User{},
		&StudentProfile{},
		&ProductOrder{},
		&Course{},
		&Class{},
		&ClassStudent{},
	)
	if err != nil {
		log.Fatal(err)
	}

	tdd := Course{
		Name:        "TDD",
		Description: "TDD is Fun",
	}
	db.Save(&tdd)

	pong := User{Username: "pong"}
	gap := User{Username: "gap"}
	pae := User{Username: "pae"}
	ball := User{Username: "ball"}
	db.Save(&pong)
	db.Save(&gap)
	db.Save(&pae)
	db.Save(&ball)

	class := Class{
		CourseID:  tdd.ID,
		TrainerID: pong.ID,
		Start:     time.Date(2023, 5, 10, 9, 0, 0, 0, time.Local),
		End:       time.Date(2023, 5, 12, 17, 0, 0, 0, time.Local),
		Seats:     10,
		Students: []ClassStudent{
			{StudentID: ball.ID},
			{StudentID: pae.ID},
			{StudentID: gap.ID},
		},
	}
	db.Save(&class)

	var foundClass Class
	db.First(&foundClass, class.ID)

	db.Preload("Course").Preload("Trainer").Preload("Students.Student").First(&foundClass, class.ID)

	fmt.Println("#ID: ", foundClass.ID)
	fmt.Println("Name: ", foundClass.Course.Name)
	fmt.Println("Description: ", foundClass.Course.Description)
	fmt.Println("\tBy: ", foundClass.Trainer.Username)
	fmt.Println("\tDate: ", foundClass.Start, foundClass.End)
	fmt.Println("Students: ")
	for _, student := range foundClass.Students {
		fmt.Println("\tName: ", student.Student.Username)
	}

}

func PrintOrder(order Order) {
	fmt.Println()
	fmt.Printf("Order ID : %v\n", order.ID)
	fmt.Printf("Products : ")
	for _, p := range order.Products {
		fmt.Printf("\t%v\t\t%v\t%v\n", p.Product.Name, p.Product.Price, p.Amount)
	}
}

// user := User{
// 	// Model: gorm.Model{
// 	// 	ID: 1,
// 	// },
// 	Username: "patna",
// 	Profile: StudentProfile{
// 		CompanyName: "PEA",
// 		JobTitle:    "Engineer",
// 		Level:       "Noob",
// 	},
// 	// Password: "patna",
// 	// Role:     "STUDENT",
// }

// db.Create(&user)

// user2 := User{
// 	Username: "patna2",
// 	Profile: StudentProfile{
// 		CompanyName: "PEA",
// 		JobTitle:    "Engineer",
// 		Level:       "Newbie",
// 	},
// 	// Password: "patna",
// 	// Role:     "STUDENT",
// }

// db.Create(&user2)

// found := db.Model(&User{}).Where("username=?", "patna")
// fmt.Printf("%+v", found)

// studentpro := StudentProfile{
// 	CompanyName: "PEA",
// 	JobTitle:    "Engineer",
// 	Level:       "noob",
// 	UserID:      user.ID,
// }

// db.Create(&studentpro)

// var found User
// db.Preload("Profile").First(&found, 1)
// fmt.Printf("%+v \n\n", found)

// shirt := Product{
// 	Name:  "T-Shirt",
// 	Price: 350,
// 	Stock: 200,
// }

// db.Create(&shirt)

// order := Order{
// 	ProductID: shirt.ID,
// 	Amount:    10,
// }

// order2 := Order{
// 	ProductID: shirt.ID,
// 	Amount:    20,
// }

// db.Create(&order)
// db.Create(&order2)

// var found Product
// db.Preload("Orders").First(&found, 1)
// fmt.Printf("%+v \n\n", found)

// var found2 Order
// db.Preload("Product").First(&found2, 1)
// fmt.Printf("%+v", found2)

// found := []Product{}
// db.Find(&found)
// fmt.Print(found)

// shirt2 := new(Product)
// db.First(&shirt2, 1)

//db.Model(&Product{}).Where("id=?", "1").Update("name", "hello")

// update := Product{
// 	Name:  "Eiei",
// 	Stock: 222,
// }
// db.Model(&Product{}).Where("id=?", "1").Updates(&update)

// db.Model(&Product{}).Where("id=?", "1").Updates(map[string]interface{}{
// 	"name":  "helloword",
// 	"stock": 123,
// })

// shirt := Product{
// 	Name:  "T-Shirt",
// 	Price: 350,
// 	Stock: 200,
// }

// short := Product{
// 	Name:  "Short v1",
// 	Price: 600,
// 	Stock: 150,
// }

// toy := Product{
// 	Name:  "Car Toy",
// 	Price: 99,
// 	Stock: 700,
// }

// db.Create(&shirt)
// db.Create(&short)
// db.Create(&toy)

// order1 := Order{
// 	Products: []ProductOrder{
// 		{ProductID: shirt.ID, Amount: 1},
// 		{ProductID: short.ID, Amount: 1},
// 	},
// }
// db.Create(&order1)

// order2 := Order{
// 	Products: []ProductOrder{
// 		{ProductID: shirt.ID, Amount: 1},
// 		{ProductID: toy.ID, Amount: 1},
// 	},
// }
// db.Create(&order2)

// var foundOrder Order
// db.Preload("Products.Product").First(&foundOrder, order1.ID)
// PrintOrder(foundOrder)
