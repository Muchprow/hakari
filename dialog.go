package hakari

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type filePayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (a *App) setupDialogBridge() {
	a.w.Bind("hakari_internal_saveFile", func(payload filePayload) string {
		url, err := a.saveToCache(payload.Name, payload.Data)
		if err != nil {
			fmt.Printf("hakari: save to cache failed: %v\n", err)
			return ""
		}
		return url
	})
}

func (a *App) saveToCache(name string, dataURL string) (string, error) {
	dir, err := a.SaveDir()
	if err != nil {
		return "", err
	}

	cacheDir := filepath.Join(dir, "cache")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", fmt.Errorf("hakari: cannot create cache dir: %w", err)
	}

	base64Data := dataURL
	if i := strings.Index(dataURL, ","); i >= 0 {
		base64Data = dataURL[i+1:]
	}

	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", fmt.Errorf("hakari: cannot decode base64: %w", err)
	}

	hash := sha256.Sum256(decoded)
	ext := filepath.Ext(name)
	if ext == "" {
		ext = ".bin"
	}
	filename := hex.EncodeToString(hash[:8]) + ext

	fullPath := filepath.Join(cacheDir, filename)

	if err := os.WriteFile(fullPath, decoded, 0644); err != nil {
		return "", fmt.Errorf("hakari: cannot write cache file: %w", err)
	}

	a.ensureAssetServer()

	if a.assets != nil {
		return a.assets.URL(filename), nil
	}

	return fullPath, nil
}

func (a *App) OpenFile() {
	a.w.Dispatch(func() {
		a.w.Eval(`(function(){
var input = document.createElement('input');
input.type = 'file';
input.style.display = 'none';
input.onchange = function() {
    var file = input.files[0];
    if (!file) return;
    var reader = new FileReader();
    reader.onload = function() {
        window.hakari_internal_saveFile({name: file.name, data: reader.result})
            .then(function(url){ window.hakariFileSelected && window.hakariFileSelected(url, file.name); })
            .catch(function(err){ console.error('hakari: file save failed', err); });
    };
    reader.readAsDataURL(file);
};
document.body.appendChild(input);
input.click();
setTimeout(function(){ document.body.removeChild(input); }, 1000);
})();`)
	})
}
