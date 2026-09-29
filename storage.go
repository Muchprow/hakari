package hakari

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func (a *App) SaveDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("hakari: cannot find config dir: %w", err)
	}

	dir := filepath.Join(base, "Hakari", a.title)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("hakari: cannot create save dir: %w", err)
	}

	return dir, nil
}

func (a *App) Save(slot string, data any) error {
	dir, err := a.SaveDir()
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("hakari: cannot marshal save data: %w", err)
	}

	path := filepath.Join(dir, slot+".json")

	if err := os.WriteFile(path, jsonData, 0644); err != nil {
		return fmt.Errorf("hakari: cannot write save file: %w", err)
	}

	return nil
}

func (a *App) Load(slot string, data any) error {
	dir, err := a.SaveDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, slot+".json")

	jsonData, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("hakari: cannot read save file: %w", err)
	}

	if err := json.Unmarshal(jsonData, data); err != nil {
		return fmt.Errorf("hakari: cannot unmarshal save data: %w", err)
	}

	return nil
}

func (a *App) SaveExists(slot string) bool {
	dir, err := a.SaveDir()
	if err != nil {
		return false
	}

	path := filepath.Join(dir, slot+".json")
	_, err = os.Stat(path)
	return err == nil
}

func (a *App) Delete(slot string) error {
	dir, err := a.SaveDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, slot+".json")

	if err := os.Remove(path); err != nil {
		return fmt.Errorf("hakari: cannot delete save file: %w", err)
	}

	return nil
}
