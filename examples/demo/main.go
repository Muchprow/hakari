package main

import (
	"embed"

	"github.com/Muchprow/hakari"
)

//go:embed assets
var assetsFS embed.FS

type GameState struct {
	PlayerName string `json:"playerName"`
	HP         int    `json:"hp"`
	Level      int    `json:"level"`
}

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
				var state GameState
				if s.SaveExists("autosave") {
					if err := s.Load("autosave", &state); err == nil {
						s.Set("playerName", state.PlayerName)
						s.Set("hp", state.HP)
						s.Set("level", state.Level)
						return
					}
				}
				s.Set("playerName", "Muchprow")
				s.Set("hp", 100)
				s.Set("level", 1)
			}
		})

		s.On("attack", func() int {
			hp := s.Get("hp").(int) - 10
			s.Set("hp", hp)
			autoSave(s)
			return hp
		})

		s.On("up", func() int {
			level := s.Get("level").(int) + 1
			s.Set("level", level)
			autoSave(s)
			return level
		})

		s.On("pickFile", func() {
			s.OpenFile()
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

func autoSave(s *hakari.Screen) {
	state := GameState{
		PlayerName: s.Get("playerName").(string),
		HP:         s.Get("hp").(int),
		Level:      s.Get("level").(int),
	}
	if err := s.Save("autosave", state); err != nil {
		println("save error:", err.Error())
	}
}
