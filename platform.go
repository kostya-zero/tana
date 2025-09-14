package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

func GetStorePath() (string, error) {
	switch runtime.GOOS {
	case "windows":
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData == "" {
			return "", errors.New("LOCALAPPDATA not set")
		}
		return filepath.Join(localAppData, "tana.tanadb"), nil
	case "darwin":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, "Library", "Application Support", "tana.tanadb"), nil
	default:
		xdg := os.Getenv("XDG_CONFIG_HOME")
		if xdg == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			xdg = filepath.Join(home, ".config")
		}
		return filepath.Join(xdg, "tana.tanadb"), nil
	}
}
