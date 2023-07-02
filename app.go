package main

import (
	"context"

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

func (a *App) shutdown(ctx context.Context) {
	database.FileStore.Save()
}

// Login sets the password for the encrypted database
func (a *App) Login(password string) string {
	return database.SetPassword(password)
}
