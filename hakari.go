package hakari

import (
	"os"

	"github.com/abemedia/go-webview"
	_ "github.com/abemedia/go-webview/embedded"
)

type App struct {
	title  string
	width  int
	height int
	w      webview.WebView
}

func New(title string) *App {
	w := webview.New(true)
	w.SetTitle(title)
	w.SetSize(1024, 768, webview.HintNone)

	return &App{
		title:  title,
		width:  1024,
		height: 768,
		w:      w,
	}
}

func (a *App) LoadHTMLFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	a.w.SetHtml(string(data))
	return nil
}

func (a *App) Run() {
	defer a.w.Destroy()
	a.w.Run()
}
