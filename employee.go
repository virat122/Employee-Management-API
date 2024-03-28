package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmpName   string `json:"empname"`
	EmpSalary string `json:"salary"`
	Email     string `json:"email"`
}