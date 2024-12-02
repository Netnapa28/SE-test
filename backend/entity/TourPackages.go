package entity

import (
	"gorm.io/gorm"
)

type TourPackages struct {
	gorm.Model
	TourName	string
	PackageCode	string
	Duration	string
}