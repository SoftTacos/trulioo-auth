package main

import (
	"context"

	c "github.com/softtacos/trulioo-auth/auth/controller"
)

const (
	dbUrlEnv = "DB_URL"
)

func main() {
	controller := c.NewAuthController()
	controller.GenerateToken(context.Background(), "", "")
}
