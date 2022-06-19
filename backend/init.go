package backend

import (
	"io/fs"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

func Initialize(assets fs.FS) {
	app := &App{}

	err := wails.Run(&options.App{
		Title:     "WailsApp",
		Width:     1024,
		Height:    768,
		Assets:    assets,
		OnStartup: app.startup,
		Bind: []any{
			app,
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
