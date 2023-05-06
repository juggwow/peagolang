package model_test

import (
	"testing"
	"time"

	"github.com/juggwow/peagolang/model"
)

func TestClass_SetSeats_withValidSeats_shouldNotReturnError(t *testing.T) {
	class := model.Class{}
	err := class.SetSeats(10)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClass_SetSeats_withInvalidSeats_shouldReturnError(t *testing.T) {
	class := model.Class{}
	err := class.SetSeats(-1)
	if err == nil {
		t.Fatal("want class.SetSeats(-1) = error, got nil")
	}
}

func TestCourse_CreateClass_withEndDateBeforeStartDate_shouldReturnError(t *testing.T) {
	course := model.Course{
		Name:        "TEST",
		Description: "TEST",
	}
	start := time.Date(2023, 5, 10, 9, 0, 0, 0, time.Local)
	end := time.Date(2023, 5, 10, 8, 0, 0, 0, time.Local)
	_, err := course.CreateClass(start, end)
	if err == nil {
		t.Fatal("want course.CreateClass(start, end) = error, got nil")
	}
}

func TestClass_AddStudent_withExceedLimit_shouldReturnError(t *testing.T) {
	course := model.Course{
		Name:        "TEST",
		Description: "TEST",
	}
	start := time.Date(2023, 5, 10, 9, 0, 0, 0, time.Local)
	end := time.Date(2023, 5, 11, 17, 0, 0, 0, time.Local)
	class, err := course.CreateClass(start, end)
	if err != nil {
		t.Fatal(err)
	}
	class.SetSeats(1)
	err = class.AddStudent(model.Student{Name: "pong"})
	if err != nil {
		t.Fatal(err)
	}
	err = class.AddStudent(model.Student{Name: "gap"})
	if err == nil {
		t.Fatal("want class.AddStudent(student) = error, got nil")
	}
}

func TestClass_AddStudent_withValidSeatNumber_shouldNotReturnError(t *testing.T) {
	course := model.Course{
		Name:        "TEST",
		Description: "TEST",
	}
	start := time.Date(2023, 5, 10, 9, 0, 0, 0, time.Local)
	end := time.Date(2023, 5, 11, 17, 0, 0, 0, time.Local)
	class, err := course.CreateClass(start, end)
	if err != nil {
		t.Fatal(err)
	}
	class.SetSeats(1)
	err = class.AddStudent(model.Student{Name: "pong"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClass_AddStudent_withExistingStudent_shouldReturnError(t *testing.T) {
	course := model.Course{
		Name:        "TEST",
		Description: "TEST",
	}
	start := time.Date(2023, 5, 10, 9, 0, 0, 0, time.Local)
	end := time.Date(2023, 5, 11, 17, 0, 0, 0, time.Local)
	class, err := course.CreateClass(start, end)
	if err != nil {
		t.Fatal(err)
	}
	class.SetSeats(2)
	err = class.AddStudent(model.Student{ID: 1, Name: "pong"})
	if err != nil {
		t.Fatal(err)
	}
	err = class.AddStudent(model.Student{ID: 1, Name: "pong"})
	if err == nil {
		t.Fatal("want class.AddStudent(existingStudent) = error, got nil")
	}
}
