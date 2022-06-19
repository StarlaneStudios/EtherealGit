package backend

import (
	"context"
	"fmt"
)

type App struct {
	ctx context.Context
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	fmt.Printf("Greeting received from frontend")

	return fmt.Sprintf("Hello %s, It's really time!", name)
}
