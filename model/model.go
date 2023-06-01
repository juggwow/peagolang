package model

import (
	"errors"
	"strings"
	"time"
)

type Class struct {
	ID       uint
	Course   Course
	Trainer  Trainer
	Start    time.Time
	End      time.Time
	Seats    int
	Students []Student
}

func (c *Class) SetSeats(seats int) error {
	if seats <= 0 {
		return errors.New("invalid seats, seats can not be zero or negative")
	}
	c.Seats = seats
	return nil
}

func (c *Class) EditClass(start time.Time, end time.Time, seats int) error {

	if end.Before(start) {
		return errors.New("invalid date, end should not be before start")
	}

	if len(c.Students) > seats {
		return errors.New("invalid seats, seats must more thane number student")
	}

	c.Start = start
	c.End = end
	c.Seats = seats

	return nil
}

func (c *Class) AddStudent(student Student) error {
	if len(c.Students) >= c.Seats {
		return errors.New("student exceed seats limit")
	}
	if err := student.Profile.IsStudent(); err != nil {
		return err
	}
	for _, stu := range c.Students {
		if stu.ID == student.ID {
			return errors.New("student is already exists")
		}
	}
	c.Students = append(c.Students, student)
	return nil
}

func (c *Class) DeleteByTrainer(trainerId uint) error {
	if c.Trainer.ID != trainerId {
		return errors.New("cannot delect class by this trainer")
	}
	// if err := student.Profile.IsStudent(); err != nil {
	// 	return err
	// }
	// for _, stu := range c.Students {
	// 	if stu.ID == student.ID {
	// 		return errors.New("student is already exists")
	// 	}
	// }
	// c.Students = append(c.Students, student)
	return nil
}

type Course struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}

func (c *Course) CreateClass(start time.Time, end time.Time) (*Class, error) {
	if end.Before(start) {
		return nil, errors.New("invalid date, end should not be before start")
	}
	cls := Class{
		Course: *c,
		Start:  start,
		End:    end,
	}
	return &cls, nil
}

type Trainer struct {
	ID      uint
	Name    string
	Profile Profile
}

type Student struct {
	ID      uint
	Name    string
	Profile Profile
}

type User struct {
	ID       uint
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	ID        uint   `json:"id"`
	Firstname string `json:"fname"`
	Lastname  string `json:"lname"`
	Role      string `json:"role"`
	Company   string `json:"company"`
	Mobileno  string `json:"mobile"`
}

func (p *Profile) IsStudent() error {
	if strings.ToUpper(p.Role) != "STUDENT" {
		err := errors.New("this profile is not a student")
		return err
	}
	p.Role = "STUDENT"
	return nil
}

func (p *Profile) IsTrainer() error {
	if strings.ToUpper(p.Role) != "TRAINER" {
		err := errors.New("this profile is not a trainer")
		return err
	}
	p.Role = "TRAINER"
	return nil
}
