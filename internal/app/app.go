package app

import (
	"github.com/kwarabei/Discoverish/internal/database"
	"github.com/kwarabei/Discoverish/internal/handlers"
)

func Run() {
	api := handlers.Mount(database.Setup())

	api.Logger.Fatal(api.Start("127.0.0.1:9001"))
}
