package main

import (
	"errors"
	"fmt"
	"os"
)

func loadContext() (*Store, string, error) {
	storePath, err := GetStorePath()
	if err != nil {
		return nil, "", errors.New("cant get path to the configuration path")
	}

	store := NewStore()
	err = LoadFromFile(storePath, store)
	if err != nil {
		return nil, "", fmt.Errorf("failed to load store from file: %w", err)
	}
	return store, storePath, nil
}

func saveStore(store *Store, path string) error {
	err := SaveToFile(store, path)
	if err != nil {
		return fmt.Errorf("failed to save store to a file: %s", err.Error())
	}

	return nil
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

	err = saveStore(store, path)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}
}

func GetCommand(key string) {
	store, _, err := loadContext()
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	value, err := store.Get(key)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	println(value)
}

func ListCommand() {
	store, _, err := loadContext()
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	keys := store.GetAll()
	if len(*keys) == 0 {
		println("no keys found")
		return
	}

	for k, v := range *keys {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func DeleteCommand(key string) {
	store, path, err := loadContext()
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	err = store.Delete(key)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	err = saveStore(store, path)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}
}

func UpdateCommand(key, newValue string) {
	store, path, err := loadContext()
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	err = store.Update(key, newValue)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	err = saveStore(store, path)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}
}

func ResetCommand() {
	store, path, err := loadContext()
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	store.Reset()

	err = saveStore(store, path)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}
}
