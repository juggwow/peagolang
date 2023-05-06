package model

type Course struct{
	ID        uint `gorm:"primaryKey"`
	Name string
	Description string
}

type Trainer struct{
	Name string
}

func (c *Class) SetSeats (seats int) error {
	if seats < 0{
		return errors.New("invalid seats, seats cant be neg")
	}
	c.Seats = seats
	return nil
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

//Business logic
//Seat limit
//

type 