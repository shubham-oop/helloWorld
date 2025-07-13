package main

import (
	"github.com/shubham-oop/helloWorld/config" // 👈 This is where 'config' comes from
	"github.com/shubham-oop/helloWorld/db"
	"github.com/shubham-oop/helloWorld/internal/router"
)

func main() {
	config.InitLogger()
	config.LoadEnv()
	db.Init()

	config.Log.Info("Server starting...")
	r := router.SetupRouter()
	r.Run(":" + config.GetEnv("PORT", "8080"))
}
