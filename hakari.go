package hakari

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/abemedia/go-webview"
	_ "github.com/abemedia/go-webview/embedded"
)

type App struct {
	title   string
	width   int
	height  int
	w       webview.WebView
	data    map[string]any
	fsys    fs.FS
	root    string
	screens map[string]*Screen
	current string
	assets  *assetServer
}

func New(title string) *App {
	w := webview.New(true)
	w.SetTitle(title)
	w.SetSize(1024, 768, webview.HintNone)

	app := &App{
		title:   title,
		width:   1024,
		height:  768,
		w:       w,
		data:    make(map[string]any),
		screens: make(map[string]*Screen),
	}

	app.setupDialogBridge()

	return app
}

func (a *App) ensureAssetServer() {
	if a.assets != nil {
		return
	}

	dir, err := a.SaveDir()
	if err != nil {
		fmt.Printf("hakari: save dir failed: %v\n", err)
		return
	}

	cacheDir := filepath.Join(dir, "cache")

	server, err := newAssetServer(cacheDir)
	if err != nil {
		fmt.Printf("hakari: asset server failed: %v\n", err)
		return
	}

	a.assets = server
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

	js := fmt.Sprintf("window.hakari = window.hakari || {}; window.hakari[%q] = %s;", key, string(jsonValue))
	a.w.Dispatch(func() {
		a.w.Eval(js)
	})
}

func (a *App) Get(key string) any {
	return a.data[key]
}

func (a *App) Screen(name string, fn func(s *Screen)) {
	s := &Screen{
		app:  a,
		name: name,
	}
	fn(s)
	a.screens[name] = s
}

func (a *App) Start(name string) {
	if _, ok := a.screens[name]; !ok {
		panic(fmt.Sprintf("hakari: screen %q not found", name))
	}

	a.ensureAssetServer()

	a.GoTo(name)

	a.w.Run()
}

func (a *App) GoTo(name string) {
	a.GoToWith(name, nil)
}

func (a *App) GoToWith(name string, args map[string]any) {
	target, ok := a.screens[name]
	if !ok {
		return
	}

	if a.current != "" {
		if cur, ok := a.screens[a.current]; ok {
			for _, fn := range cur.onLeave {
				fn()
			}
		}
	}

	html, err := target.render()
	if err != nil {
		fmt.Printf("hakari: render screen %q: %v\n", name, err)
		return
	}

	htmlJSON, _ := json.Marshal(html)
	argsJSON := target.marshalArgs(args)

	script := fmt.Sprintf(`(function(){
var container = document.getElementById('hakari-screen-container');
if (!container) {
    container = document.createElement('div');
    container.id = 'hakari-screen-container';
    document.body.appendChild(container);
}

var allScreens = container.querySelectorAll('.hakari-screen');
for (var j = 0; j < allScreens.length; j++) {
    allScreens[j].style.display = 'none';
}

var existing = document.getElementById('hakari-screen-%s');
if (existing) {
    existing.style.display = 'block';
} else {
    var wrapper = document.createElement('div');
    wrapper.innerHTML = %s;
    var screenEl = wrapper.firstElementChild;
    screenEl.style.display = 'block';
    container.appendChild(screenEl);

    var scripts = screenEl.querySelectorAll('script');
    for (var i = 0; i < scripts.length; i++) {
        var oldScript = scripts[i];
        var newScript = document.createElement('script');
        newScript.textContent = oldScript.textContent;
        document.head.appendChild(newScript);
        oldScript.parentNode.removeChild(oldScript);
    }
}
window.hakariArgs = %s;
})()`, name, string(htmlJSON), argsJSON)

	a.w.Dispatch(func() {
		a.w.Eval(script)
	})

	for _, fn := range target.onEnter {
		fn()
	}

	a.current = name
}

func (a *App) Run() {
	defer a.w.Destroy()
	defer func() {
		if a.assets != nil {
			a.assets.Stop()
		}
	}()
	a.w.Run()
}
