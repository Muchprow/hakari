package main

import (
	"embed"

	"github.com/Muchprow/hakari"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	app := hakari.New("Hakari Demo")

	app.Screen("menu", func(s *hakari.Screen) {
		s.HTML("assets/screens/menu/menu.html")

		s.On("start", func() {
			s.GoTo("game")
		})

		s.On("settings", func() {
			println("settings clicked")
		})
	})

	app.Screen("game", func(s *hakari.Screen) {
		s.HTML("assets/screens/game/game.html")

		s.OnEnter(func() {
			if s.Get("hp") == nil {
				s.Set("playerName", "Muchprow")
				s.Set("hp", 100)
				s.Set("level", 1)
			}
		})

		s.On("attack", func() int {
			hp := s.Get("hp").(int) - 10
			s.Set("hp", hp)
			return hp
		})

		s.On("up", func() int {
			level := s.Get("level").(int) + 1
			s.Set("level", level)
			return level
		})

		s.On("back", func() {
			s.GoTo("menu")
		})
	})

	if err := app.LoadHTMLFromFS(assetsFS, "assets/index.html"); err != nil {
		panic(err)
	}

	app.Start("menu")
}
