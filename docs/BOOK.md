# Hakari

Lightweight Go framework for building UI-centric desktop apps and games with HTML, CSS and JavaScript.

No CGo. No Node.js. One binary.

**With love from Ukraine** 🇺🇦

---

## Table of Contents

1. [What is Hakari](#chapter-1-what-is-hakari)
2. [Installation](#chapter-2-installation)
3. [Screens](#chapter-3-screens)
4. [Data Binding](#chapter-4-data-binding)
5. [Save / Load](#chapter-5-save--load)
6. [File Dialogs](#chapter-6-file-dialogs)
7. [Asset Bundling](#chapter-7-asset-bundling)
8. [Building a Release](#chapter-8-building-a-release)
9. [Examples](#chapter-9-examples)

---

## Chapter 1. What is Hakari

**Hakari** is a lightweight framework for building desktop applications and games in **Go**, with a UI written in **HTML, CSS, and JavaScript**. No CGo. No Node.js. One binary.

### The Problem

Go is a great language for backends, utilities, and CLI tools. But when it comes to **desktop apps with a proper interface** — everything falls apart.

Existing solutions:

- **Fyne** — pure Go, but the API is awkward, thread safety is an illusion, and the rendering is far from native.
- **Wails** — beautiful UI, but pulls in **Node.js** and **WebView2 SDK**, build process is complicated, and the binary is bloated.
- **Gio** — fast, but immediate-mode, steep learning curve, and not for everyone.
- **Electron** — not even Go. 150+ MB for an empty window.

Each of them solves **part** of the problem but creates **a new one**. What you want is: **download → write HTML → write Go → run → it works**. No dependencies, no bundlers, no dancing with a tambourine.

**That didn't exist. So I built Hakari.**

### The Idea

The idea didn't come **all at once**. I tried several Go GUI frameworks. Each had its own problems — awkward APIs, poor performance, heavy dependencies.

Then I stumbled upon **`webview`** — a library that provides a **native WebView** without CGo, using prebuilt libraries. It **just worked**. And then I realized: **Go can handle the logic, HTML can handle the interface, and Hakari can connect them**.

That's how the idea was born: **a framework that doesn't fight WebView but wraps it in a human-friendly API**.

The name **Hakari** comes from a character in *Jujutsu Kaisen* who throws jackpots. It fits perfectly: a user is looking for a decent, **normal** framework — and then a jackpot hits them: **"JACKPOT! Here's Hakari, my framework."**

### About Me

My name is **Mykhailo** (Michael). I'm from **Ukraine** 🇺🇦, and I'm **15 years old**. I write in **C, C++, C#, Go, Python, and Java** — I don't lock myself into one language; I pick whatever fits the task.

I'm a **Java enjoyer**. That might sound strange for a Go project — but Go caught my attention for the same reasons I love Java: **native compilation for any platform**, **even resource distribution**, and **predictable performance**. Go takes the best of Java's philosophy and pushes it further.

Before Hakari, I built several projects: **Wincess** — a Windows utility in C++ (bypassing activation watermarks and enabling personalization), **Owl-UI** — a Python GUI framework, **Live-Reload** — a C++ DLL hot-swapper, and **many other things I studied and tested on myself**.

I'm not a "professional developer." I'm a **curious one**. I learn by building. I break things to understand how they work. And I build tools I want to use myself.

### Philosophy

Hakari is built on **five principles**:

1. **Simplicity.** Download → run → it works. No configs, no bundlers, no Node.js.
2. **One binary.** All assets are embedded. Copy the `.exe` — everything works.
3. **Transparency.** DevTools like in a browser. Debug your UI as a web page.
4. **Versatility.** Not just games. Any UI-centric applications.
5. **Performance.** Under 100 MB of memory. Fast builds. No dependencies.

**Hakari is not "just another GUI framework."** It's a **philosophy**: Go for logic, HTML for the interface, Hakari for the connection. Everything else is up to you.

---

## Chapter 2. Installation

### Requirements

- **Go 1.21+** — [download here](https://go.dev/dl/)
- A code editor — VS Code, GoLand, or anything you like
- **Windows**: WebView2 Runtime (usually pre-installed on Windows 10/11)

### Install Hakari

In your project folder:

```bash
go mod init myapp
go get github.com/Muchprow/hakari
```

### First project

Create `main.go`:

```go
package main

import (
	"embed"

	"github.com/Muchprow/hakari"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	app := hakari.New("My App")

	app.Screen("main", func(s *hakari.Screen) {
		s.HTML("assets/screens/main/main.html")
	})

	if err := app.LoadHTMLFromFS(assetsFS, "assets/index.html"); err != nil {
		panic(err)
	}

	app.Start("main")
}
```

### Project structure

```
myapp/
├── main.go
├── go.mod
├── go.sum
└── assets/
    ├── index.html
    ├── index.css
    ├── index.js
    └── screens/
        └── main/
            ├── main.html
            ├── main.css
            └── main.js
```

### `assets/index.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My App</title>
    <link rel="stylesheet" href="index.css">
</head>
<body>
    <div id="hakari-screen-container"></div>
    <script src="index.js"></script>
</body>
</html>
```

### `assets/screens/main/main.html`

```html
<h1>Hello, Hakari!</h1>
```

### Run it

```bash
go run main.go
```

A window opens with **"Hello, Hakari!"** — that's it. No build step, no Node.js, no configs.

### Build a binary

```bash
go build -ldflags="-H windowsgui" -o myapp.exe
```

`-H windowsgui` hides the console window on Windows. The `.exe` is self-contained — copy it anywhere, it works.

---

## Chapter 3. Screens

Every screen has its own **HTML, CSS, and JS**. Screens are registered in Go and switch via `GoTo`. State is preserved between transitions.

### Register a screen

```go
app.Screen("menu", func(s *hakari.Screen) {
	s.HTML("assets/screens/menu/menu.html")
})

app.Screen("game", func(s *hakari.Screen) {
	s.HTML("assets/screens/game/game.html")
})
```

### Start with a screen

```go
app.Start("menu")
```

### Navigate between screens

```go
app.Screen("menu", func(s *hakari.Screen) {
	s.HTML("assets/screens/menu/menu.html")

	s.On("start", func() {
		s.GoTo("game")
	})
})
```

In HTML:

```html
<button onclick="window.menu_start()">Start Game</button>
```

### Lifecycle hooks

```go
s.OnEnter(func() {
	// runs every time the screen is shown
})

s.OnLeave(func() {
	// runs every time the screen is hidden
})
```

### Why state is preserved

When you navigate away, the screen's DOM is **hidden**, not destroyed. Timers, JS variables, and input values survive. When you come back, everything is as you left it.

---

## Chapter 4. Data Binding

Go and JavaScript share data through `window.hakari`. Use `Set` to send data, `Get` to read it, and `On` to handle events.

### Set data from Go

```go
s.Set("hp", 100)
s.Set("playerName", "Muchprow")
```

JavaScript reads:

```js
const hp = window.hakari['game.hp']
const name = window.hakari['game.playerName']
```

Note: data is namespaced by screen name. `s.Set("hp", 100)` on screen `game` becomes `game.hp` in JS.

### Read data in Go

```go
hp := s.Get("hp").(int)
```

### Handle events from JS

```go
s.On("attack", func(dmg int) int {
	hp := s.Get("hp").(int) - dmg
	s.Set("hp", hp)
	return hp
})
```

In HTML:

```html
<button onclick="window.game_attack(10)">Attack</button>
```

In JS:

```js
await window.game_attack(10)
```

### Global data

Data set on `app` (not `s`) is available under its own key:

```go
app.Set("version", "1.0.0")
```

```js
window.hakari['version']
```

---

## Chapter 5. Save / Load

Save and load JSON data to the OS-standard config directory.

### Where files go

- **Windows:** `%APPDATA%\Hakari\<app-name>\`
- **Linux:** `~/.config/Hakari/<app-name>/`
- **macOS:** `~/Library/Application Support/Hakari/<app-name>/`

### Define a save structure

```go
type GameState struct {
	PlayerName string `json:"playerName"`
	HP         int    `json:"hp"`
	Level      int    `json:"level"`
}
```

### Save

```go
state := GameState{
	PlayerName: "Muchprow",
	HP:         60,
	Level:      7,
}

if err := s.Save("autosave", state); err != nil {
	// handle error
}
```

### Load

```go
var state GameState
if err := s.Load("autosave", &state); err != nil {
	// handle error
}
```

### Check if a save exists

```go
if s.SaveExists("autosave") {
	// load it
}
```

### Delete

```go
s.Delete("autosave")
```

### Autosave pattern

```go
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
		s.Set("playerName", "Player")
		s.Set("hp", 100)
		s.Set("level", 1)
	}
})
```

---

## Chapter 6. File Dialogs

Open any file from disk — audio, video, images, documents — and use it in the UI.

### Open a file

```go
s.On("pickFile", func() {
	s.OpenFile()
})
```

In HTML:

```html
<button onclick="window.game_pickFile()">Open File</button>
```

### Handle the selected file

The user picks a file, Go saves it to cache, and returns a local URL. Hook into `window.hakariFileSelected`:

```js
window.hakariFileSelected = function (url, name) {
	if (!url) {
		console.error('hakari: file load failed')
		return
	}
	console.log('selected:', name, url)

	var player = document.getElementById('player')
	if (player) {
		player.src = url
	}
}
```

### How it works

1. Go opens a hidden `<input type="file">`.
2. User picks a file.
3. JS reads it as base64, sends to Go.
4. Go saves it to `SaveDir/cache/` with a SHA256-based name.
5. Go returns a URL like `http://127.0.0.1:54321/<hash>.mp3`.
6. JS uses that URL in `<audio>`, `<video>`, or `<img>`.

### Why a local HTTP server

Browsers block `file://` URLs for security. A local HTTP server is the standard workaround. It listens on `127.0.0.1` on a random port, serves only files from the cache directory, and protects against path traversal.

---

## Chapter 7. Asset Bundling

Hakari inlines CSS and JS into a single HTML document at load time. You write separate files; the framework combines them.

### What you write

```
assets/screens/game/
├── game.html
├── game.css
└── game.js
```

### `game.html`

```html
<link rel="stylesheet" href="game.css">

<h1>Game</h1>
<button onclick="window.game_attack()">Attack</button>

<script src="game.js"></script>
```

### What the user gets

A single HTML document with `<style>` and `<script>` blocks inlined. No extra requests, no `file://` paths, no CORS issues.

### Why this matters

- **One binary.** All assets are embedded via `//go:embed`.
- **No server.** No static file serving, no path resolution.
- **Familiar workflow.** Write HTML, CSS, JS like on the web.

### Path rules

Asset references inside a screen are resolved relative to the screen's own folder. So `game.css` in `game.html` means `assets/screens/game/game.css`. No need for long relative paths.

---

## Chapter 8. Building a Release

### Standard build

```bash
go build -o myapp.exe
```

This produces a working binary with a console window attached.

### GUI build (no console)

```bash
go build -ldflags="-H windowsgui" -o myapp.exe
```

`-H windowsgui` hides the console window on Windows. Recommended for release.

### Cross-platform build

Hakari uses **no CGo**, so cross-compilation works out of the box:

```bash
GOOS=windows GOARCH=amd64 go build -o myapp.exe
GOOS=linux   GOARCH=amd64 go build -o myapp
GOOS=darwin  GOARCH=amd64 go build -o myapp
```

### What's inside the binary

Everything. `//go:embed` packs all assets into the executable. Copy the `.exe` to another machine — it works without any external files.

### Size and memory

- **Binary size:** typically 5-15 MB
- **Memory usage:** under 100 MB for typical apps

---

## Chapter 9. Examples

### Todo list

A simple app with add, toggle, and delete.

```go
type Todo struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

app.Screen("main", func(s *hakari.Screen) {
	s.HTML("assets/screens/main/main.html")

	s.On("add", func(text string) []Todo {
		todos := getTodos(s)
		todos = append(todos, Todo{Text: text})
		s.Set("todos", todos)
		return todos
	})

	s.On("toggle", func(index int) []Todo {
		todos := getTodos(s)
		todos[index].Done = !todos[index].Done
		s.Set("todos", todos)
		return todos
	})

	s.On("delete", func(index int) []Todo {
		todos := getTodos(s)
		todos = append(todos[:index], todos[index+1:]...)
		s.Set("todos", todos)
		return todos
	})
})

func getTodos(s *hakari.Screen) []Todo {
	raw := s.Get("todos")
	if raw == nil {
		return []Todo{}
	}
	return raw.([]Todo)
}
```

### Wrapper around a website

Open an external site inside a Hakari window.

```go
app.Screen("main", func(s *hakari.Screen) {
	s.HTML("assets/screens/main/main.html")
})

s.On("open", func() {
	s.GoTo("web")
})

app.Screen("web", func(s *hakari.Screen) {
	s.HTML("assets/screens/web/web.html")
})
```

In `web.html`:

```html
<iframe src="https://example.com" style="width:100%;height:100vh;border:none;"></iframe>
```

### Simple game

A minimal clicker.

```go
app.Screen("game", func(s *hakari.Screen) {
	s.HTML("assets/screens/game/game.html")

	s.OnEnter(func() {
		if s.Get("score") == nil {
			s.Set("score", 0)
		}
	})

	s.On("click", func() int {
		score := s.Get("score").(int) + 1
		s.Set("score", score)
		return score
	})
})
```

In `game.html`:

```html
<h1>Score: <span id="score">0</span></h1>
<button onclick="window.game_click()">Click</button>

<script>
document.addEventListener('click', function () {
	var score = window.hakari['game.score']
	var el = document.getElementById('score')
	if (el && score !== undefined) el.textContent = score
})
</script>
```

### More examples

Full working examples live in the [`examples/`](../examples) folder of the repository.