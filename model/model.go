package model

import (
	"errors"
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

func (c *Class) AddStudent(student Student) error {
	if len(c.Students) >= c.Seats {
		return errors.New("student exceed seats limit")
	}
	for _, stu := range c.Students {
		if stu.ID == student.ID {
			return errors.New("student is already exists")
		}
	}
	c.Students = append(c.Students, student)
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
	ID   uint
	Name string
}

type Student struct {
	ID   uint
	Name string
}

type User struct {
	ID       uint
	Username string
	Password string
}
