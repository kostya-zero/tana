package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

func GetConfigDir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData == "" {
			return "", errors.New("LOCALAPPDATA not set")
		}
		return filepath.Join(localAppData, "tana"), nil
	case "darwin":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, "Library", "Application Support", "tana"), nil
	default:
		xdg := os.Getenv("XDG_CONFIG_HOME")
		if xdg == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			xdg = filepath.Join(home, ".config")
		}
		return filepath.Join(xdg, "tana"), nil
	}
}
