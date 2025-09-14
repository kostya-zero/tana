package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Store struct {
	Store map[string]string
}

func NewStore() *Store {
	return &Store{Store: make(map[string]string)}
}

func SaveToFile(store *Store, path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)
	err = encoder.Encode(store.Store)
	if err != nil {
		return err
	}

	return nil
}

func LoadFromFile(path string, store *Store) error {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}
	if info.Size() == 0 {
		return nil
	}

	decoder := gob.NewDecoder(f)
	store.Store = make(map[string]string)
	err = decoder.Decode(&store.Store)
	if err != nil {
		switch {
		case err == io.EOF:
			return nil
		case errors.Is(err, io.ErrUnexpectedEOF):
			return fmt.Errorf("file %s is corrupted or incomplete", path)
		}
		return err
	}

	return nil
}

func (s *Store) Set(key, value string) error {
	_, ok := s.Store[key]
	if ok {
		return fmt.Errorf("key `%s` is already exist", key)
	}

	s.Store[key] = value
	return nil
}

func (s *Store) Get(key string) (string, error) {
	value, ok := s.Store[key]
	if !ok {
		return "", fmt.Errorf("key `%s` is not found", key)
	}
	return value, nil
}

func (s *Store) Update(key, newValue string) error {
	_, ok := s.Store[key]
	if !ok {
		return fmt.Errorf("key `%s` is not found", key)
	}

	s.Store[key] = newValue
	return nil
}

func (s *Store) Delete(key string) error {
	_, ok := s.Store[key]
	if !ok {
		return fmt.Errorf("key `%s` is not found", key)
	}

	delete(s.Store, key)
	return nil
}

func (s *Store) GetAll() *map[string]string {
	return &s.Store
}

func (s *Store) Reset() {
	s.Store = make(map[string]string)
}
