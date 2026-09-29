# Hakari

Lightweight Go framework for building UI-centric desktop apps and games with HTML, CSS and JavaScript.

## With love from Ukraine💙💛

## Features

- **No CGo** — pure Go, works everywhere Go works
- **No Node.js** — no build step, no npm
- **One binary** — assets embedded via `go:embed`
- **HTML/CSS/JS UI** — write your interface like a web page
- **Multi-screen apps** — each screen has its own HTML/CSS/JS
- **State preservation** — screen state survives navigation
- **Reactive data binding** — Go and JS share data seamlessly
- **Event system** — call Go from JavaScript, update JS from Go
- **Save/load** — JSON persistence in OS-standard config directory
- **Built-in DevTools** — debug your UI like a web page
- **Low memory footprint** — under 100 MB for typical apps
- **Asset bundling** — separate CSS/JS files are inlined automatically

## Status

Early development.

## Requirements

- Go 1.21+

## Install

go get github.com/Muchprow/hakari


## Quick start

```go
package main

import "github.com/Muchprow/hakari"

func main() {
	app := hakari.New("My App")
	app.LoadHTMLFile("assets/index.html")
	app.Run()
}
