package main

import (
	"encoding/gob"
	"os"
)

type Store struct {
	store map[string]string
}

func NewStore() *Store {
	return &Store{store: make(map[string]string)}
}

func SaveToFile(store *Store, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)
	err = encoder.Encode(store)
	if err != nil {
		return err
	}

	return nil
}

func LoadFromFile(path string, store *Store) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := gob.NewDecoder(f)
	err = decoder.Decode(store)
	if err != nil {
		return err
	}

	return nil
}
