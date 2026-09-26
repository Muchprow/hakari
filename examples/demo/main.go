package main

import (
	"embed"

	"github.com/Muchprow/hakari"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	app := hakari.New("Hakari Demo")

	app.On("onReady", func() {
		app.Set("playerName", "Muchprow")
		app.Set("hp", 100)
		app.Set("level", 1)
	})

	app.On("takeDamage", func(damage int) int {
		current := app.Get("hp").(int)
		current -= damage
		app.Set("hp", current)
		return current
	})

	app.On("levelUp", func() int {
		level := app.Get("level").(int)
		level++
		app.Set("level", level)
		return level
	})

	if err := app.LoadHTMLFromFS(assetsFS, "assets/index.html"); err != nil {
		panic(err)
	}
	app.Run()
}
