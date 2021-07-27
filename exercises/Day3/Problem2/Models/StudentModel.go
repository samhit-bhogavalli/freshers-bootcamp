package Models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string
	LastName  string
	DOB       string
	Address   string
	Marks     []Marks
}

func (s *Student) TableName() string {
	return "student"
}

type Marks struct {
	gorm.Model
	StudentID uint
	Subject   string
	Score     int
}

func (m *Marks) TableName() string {
	return "marks"
}
