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

	// AppMenu := menu.NewMenu()
	// FileMenu := AppMenu.AddSubmenu("File")
	// FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
	// 	fmt.Println("Open")
	// })
	// FileMenu.AddSeparator()
	// FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
	// 	wails_runtime.Quit(app.ctx)
	// })

	// if runtime.GOOS == "darwin" {
	// 	AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	// 	AppMenu.Append(menu.AppMenu())  // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	// }

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Mech Commander",
		Width:  1024,
		Height: 768,
		// Menu:             AppMenu,
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
