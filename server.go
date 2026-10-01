package hakari

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type assetServer struct {
	server   *http.Server
	port     int
	cacheDir string
}

func newAssetServer(cacheDir string) (*assetServer, error) {
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("hakari: cannot create cache dir: %w", err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("hakari: cannot start asset server: %w", err)
	}

	port := listener.Addr().(*net.TCPAddr).Port

	s := &assetServer{
		port:     port,
		cacheDir: cacheDir,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handle)

	s.server = &http.Server{Handler: mux}

	go func() {
		s.server.Serve(listener)
	}()

	return s, nil
}

func (s *assetServer) handle(w http.ResponseWriter, r *http.Request) {
	requested := strings.TrimPrefix(r.URL.Path, "/")
	if requested == "" {
		http.NotFound(w, r)
		return
	}

	fullPath, err := filepath.Abs(filepath.Join(s.cacheDir, requested))
	if err != nil {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}

	cacheAbs, _ := filepath.Abs(s.cacheDir)
	if !strings.HasPrefix(fullPath, cacheAbs) {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	http.ServeFile(w, r, fullPath)
}

func (s *assetServer) URL(filename string) string {
	return fmt.Sprintf("http://127.0.0.1:%d/%s", s.port, filename)
}

func (s *assetServer) Stop() {
	if s.server != nil {
		s.server.Close()
	}
}
