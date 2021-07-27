package Models

import (
	"gorm.io/gorm/clause"
	"local/exercises/Day3/Problem2/Config"
)
import _ "gorm.io/driver/mysql"

func GetAllStudents(student *[]Student) error {
	if er := Config.DB.Preload(clause.Associations).Find(student).Error; er != nil {
		return er
	}
	return nil
}

func CreateStudent(student *Student) error {
	if er := Config.DB.Create(student).Error; er != nil {
		return er
	}
	return nil
}

func GetStudentById(student *Student, id string) error {
	if er := Config.DB.Preload(clause.Associations).Where("id = ?", id).First(student).Error; er != nil {
		return er
	}
	return nil
}

func UpdateStudent(student *Student, id string) error {
	if er := Config.DB.Save(student).Error; er != nil {
		return er
	}
	return nil
}

func DeleteStudent(student *Student, id string) error {
	if er := Config.DB.Where("id = ?", id).Delete(student).Error; er != nil {
		return er
	}
	return nil
}
