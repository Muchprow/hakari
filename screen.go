package hakari

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
	"regexp"
	"strings"
)

type Screen struct {
	app     *App
	name    string
	html    string
	onEnter []func()
	onLeave []func()
	loaded  bool
}

func (s *Screen) HTML(filePath string) {
	s.html = filePath
}

func (s *Screen) On(event string, handler any) {
	s.app.w.Bind(s.name+"_"+event, handler)
}

func (s *Screen) Set(key string, value any) {
	s.app.Set(s.name+"."+key, value)
}

func (s *Screen) Get(key string) any {
	return s.app.Get(s.name + "." + key)
}

func (s *Screen) GoTo(name string) {
	s.app.GoTo(name)
}

func (s *Screen) GoToWith(name string, args map[string]any) {
	s.app.GoToWith(name, args)
}

func (s *Screen) OnEnter(fn func()) {
	s.onEnter = append(s.onEnter, fn)
}

func (s *Screen) OnLeave(fn func()) {
	s.onLeave = append(s.onLeave, fn)
}

var (
	cssLinkRe2 = regexp.MustCompile(`<link[^>]*rel="stylesheet"[^>]*href="([^"]+)"[^>]*>`)
	jsSrcRe2   = regexp.MustCompile(`<script[^>]*src="([^"]+)"[^>]*></script>`)
)

func (s *Screen) render() (string, error) {
	if s.html == "" {
		return "", fmt.Errorf("hakari: screen %q has no HTML file", s.name)
	}

	htmlPath := s.html
	dir := path.Dir(htmlPath)
	base := path.Base(htmlPath)

	htmlData, err := fs.ReadFile(s.app.fsys, htmlPath)
	if err != nil {
		return "", fmt.Errorf("hakari: read screen %q: %w", s.name, err)
	}

	html := string(htmlData)
	html = s.app.inlineInDir(html, dir, base)
	html = s.wrapScreenHTML(html)

	s.loaded = true
	return html, nil
}

func (s *Screen) wrapScreenHTML(inner string) string {
	return fmt.Sprintf(`<div class="hakari-screen" id="hakari-screen-%s" data-screen="%s">%s</div>`,
		s.name, s.name, inner)
}

func (s *Screen) wrapScreenJS(js string) string {
	return fmt.Sprintf(`(function(){
var __screen = %q;
var __hakari = window.hakari || {};
window.hakariScreen = function(key, fallback) {
    var v = __hakari[__screen + "." + key];
    return v !== undefined ? v : fallback;
};
window.hakariSet = function(key, value) {
    __hakari[__screen + "." + key] = value;
    if (window.renderScreens) window.renderScreens();
};
%s
})();`, s.name, js)
}

func (s *Screen) marshalArgs(args map[string]any) string {
	if args == nil {
		return "{}"
	}
	data, err := json.Marshal(args)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func (a *App) inlineInDir(html string, dir string, base string) string {
	html = cssLinkRe2.ReplaceAllStringFunc(html, func(match string) string {
		sub := cssLinkRe2.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		href := sub[1]
		if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
			return match
		}
		cssPath := path.Join(dir, path.Base(href))
		data, err := fs.ReadFile(a.fsys, cssPath)
		if err != nil {
			return match
		}
		return "<style>\n" + string(data) + "\n</style>"
	})

	html = jsSrcRe2.ReplaceAllStringFunc(html, func(match string) string {
		sub := jsSrcRe2.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		src := sub[1]
		if strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
			return match
		}
		jsPath := path.Join(dir, path.Base(src))
		data, err := fs.ReadFile(a.fsys, jsPath)
		if err != nil {
			return match
		}
		wrapped := wrapJS(string(data), path.Base(src))
		return "<script>\n" + wrapped + "\n</script>"
	})

	return html
}

func wrapJS(js string, name string) string {
	return fmt.Sprintf(`(function(){
%s
})();`, js)
}
