# Hakari

Lightweight Go framework for building UI-centric desktop apps and games with HTML, CSS and JavaScript.

No CGo. No Node.js. One binary.

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