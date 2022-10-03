package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

type Mech struct {
	Id int
	X  int
	Y  int
}

type GameApi struct {
	Mechs []Mech
}

func (g *GameApi) Move(id, x, y int) bool {
	for _, m := range g.Mechs {
		if m.Id == id {
			m.X = x
			m.Y = y

			return true
		}
	}

	return false
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Mech Commander",
		Width:            730,
		Height:           770,
		DisableResize:    true,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
