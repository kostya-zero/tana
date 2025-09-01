package main

import (
	"os"
	"path/filepath"
)

func SetCommand(key, value string) {
	configDir, err := GetConfigDir()
	if err != nil {
		PrintError("Cant get path to the configuration path")
		os.Exit(1)
	}

	storePath := filepath.Join(configDir, "data.tanadb")

	store := NewStore()
	err = LoadFromFile(storePath, store)
	if err != nil {
		PrintError("failed to load store from file: " + err.Error())
		os.Exit(1)
	}

	err = store.Set(key, value)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	err = SaveToFile(store, storePath)
	if err != nil {
		PrintError("failed to save store to a file" + err.Error())
		os.Exit(1)
	}
}
