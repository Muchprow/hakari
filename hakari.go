package hakari

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"

	"github.com/abemedia/go-webview"
	_ "github.com/abemedia/go-webview/embedded"
)

type App struct {
	title  string
	width  int
	height int
	w      webview.WebView
	data   map[string]any
	fsys   fs.FS
	root   string
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
		data:   make(map[string]any),
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

func (a *App) LoadHTMLFromFS(fsys fs.FS, path string) error {
	a.fsys = fsys
	a.root = strings.TrimSuffix(path, "/index.html")

	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return err
	}

	html := a.inlineAssets(string(data))
	a.w.SetHtml(html)
	return nil
}

var (
	cssLinkRe = regexp.MustCompile(`<link[^>]*rel="stylesheet"[^>]*href="([^"]+)"[^>]*>`)
	jsSrcRe   = regexp.MustCompile(`<script[^>]*src="([^"]+)"[^>]*></script>`)
)

func (a *App) inlineAssets(html string) string {
	html = cssLinkRe.ReplaceAllStringFunc(html, func(match string) string {
		sub := cssLinkRe.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		href := sub[1]
		path := a.root + "/" + strings.TrimPrefix(href, "./")
		data, err := fs.ReadFile(a.fsys, path)
		if err != nil {
			return match
		}
		return "<style>\n" + string(data) + "\n</style>"
	})

	html = jsSrcRe.ReplaceAllStringFunc(html, func(match string) string {
		sub := jsSrcRe.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		src := sub[1]
		path := a.root + "/" + strings.TrimPrefix(src, "./")
		data, err := fs.ReadFile(a.fsys, path)
		if err != nil {
			return match
		}
		return "<script>\n" + string(data) + "\n</script>"
	})

	return html
}

func (a *App) On(name string, handler any) {
	a.w.Bind(name, handler)
}

func (a *App) Set(key string, value any) {
	a.data[key] = value

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return
	}

	js := fmt.Sprintf("window.hakari = window.hakari || {}; window.hakari.%s = %s;", key, string(jsonValue))
	a.w.Dispatch(func() {
		a.w.Eval(js)
	})
}

func (a *App) Get(key string) any {
	return a.data[key]
}

func (a *App) Run() {
	defer a.w.Destroy()
	a.w.Run()
}
