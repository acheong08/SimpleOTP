package main

import (
	"context"
	"fmt"

	"github.com/acheong08/SimpleOTP/internal/database"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Login sets the password for the encrypted database
func (a *App) Login(password string) string {
	database.SetPassword(password)
	return "success"
}
