package main

import (
	"embed"

	"github.com/Muchprow/hakari"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	app := hakari.New("Hakari Demo")

	app.On("greet", func() string {
		return "Hello from Go!"
	})

	app.On("add", func(a, b int) int {
		return a + b
	})

	if err := app.LoadHTMLFromFS(assetsFS, "assets/index.html"); err != nil {
		panic(err)
	}
	app.Run()
}
