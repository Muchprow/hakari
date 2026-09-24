package main

import (
	"embed"

	"github.com/Muchprow/hakari"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	app := hakari.New("Hakari Demo")
	if err := app.LoadHTMLFromFS(assetsFS, "assets/index.html"); err != nil {
		panic(err)
	}
	app.Run()
}
