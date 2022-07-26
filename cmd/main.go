package main

import (
	"github.com/kwarabei/Discoverish/internal/app"
)

// @title Discoverish API
// @version 1.0
// @description REST API for Discoverish application
func main() {
	// db := database.Setup()
	// db.AutoMigrate(&models.User{}, &models.Project{}, &models.Opening{})
	// db.SetupJoinTable(&models.Project{}, "Members", &models.ProjectPerson{})

	app.Run()
}
