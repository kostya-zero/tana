package main

import (
	"encoding/gob"
	"errors"
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

func (s *Store) Set(key, value string) error {
	_, ok := s.store[key]
	if ok {
		return errors.New("key already exists")
	}

	s.store[key] = value
	return nil
}

func (s *Store) Get(key string) (string, error) {
	value, ok := s.store[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}

func (s *Store) Update(key, newValue string) error {
	_, ok := s.store[key]
	if !ok {
		return errors.New("key not found")
	}

	s.store[key] = newValue
	return nil
}

func (s *Store) Delete(key string) error {
	_, ok := s.store[key]
	if !ok {
		return errors.New("key not found")
	}

	delete(s.store, key)
	return nil
}

func (s *Store) GetAll() *map[string]string {
	return &s.store
}
