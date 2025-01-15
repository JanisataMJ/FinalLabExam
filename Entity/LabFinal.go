package entity

import (
	"gorm.io/gorm"
) 

type LabFinal struct {
	gorm.Model
	SID 	string 	`valid:"required~SID is not required"`
	Name 	string	`valid:"required~Name is not required"`
	No 		uint	`valid:"required~No is not required"`
}