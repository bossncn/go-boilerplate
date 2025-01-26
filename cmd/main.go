package main

import (
	"fmt"
	"github.com/bossncn/go-boilerplate/cmd/app"
	"github.com/bossncn/go-boilerplate/config"
)

// @title Go Boilerplate
// @version 1.0
// @description Template for init projects
// @termsOfService http://swagger.io/terms/
//
// @license.name MIT
// @license.url https://github.com/bossncn/go-boilerplate/blob/main/LICENSE
func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		fmt.Println("Error Load config:", err)
		return
	}

	app.Run(cfg)
}
