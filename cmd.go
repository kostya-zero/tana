package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func loadContext() (*Store, string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return nil, "", errors.New("cant get path to the configuration path")
	}

	storePath := filepath.Join(configDir, "data.tanadb")

	store := NewStore()
	err = LoadFromFile(storePath, store)
	if err != nil {
		return nil, "", fmt.Errorf("failed to load store from file: %w", err)
	}

	return store, storePath, nil
}

func SetCommand(key, value string) {
	store, path, err := loadContext()
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	err = store.Set(key, value)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	err = SaveToFile(store, path)
	if err != nil {
		PrintError("failed to save store to a file" + err.Error())
		os.Exit(1)
	}
}
