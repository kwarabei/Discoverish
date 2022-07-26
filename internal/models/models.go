package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username      string `gorm:"unique"`
	FirstName     string
	LastName      string
	Position      string
	Qualification string
	About         string
	Skillset      pq.StringArray `gorm:"type:varchar[]"`
}

type Project struct {
	gorm.Model

	Title       string
	Description string `gorm:"type:text"`
	Website     string
	Opensource  string
	Logo        string
	PromoPic    string
	Stack       pq.StringArray `gorm:"type:varchar[]"`
	Openings    []Opening
	Members     []User `gorm:"many2many:project_people"`
}

type Opening struct {
	gorm.Model

	ProjectID uint

	Title         string
	Description   string
	Qualification string
}

type ProjectPerson struct {
	UserID    uint `gorm:"primaryKey"`
	ProjectID uint `gorm:"primaryKey"`
	IsActive  bool `gorm:"not null"`
}
