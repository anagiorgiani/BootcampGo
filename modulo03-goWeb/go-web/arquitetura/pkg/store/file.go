package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
)

func New(store Type, filename string) Store {
	switch store {
	case FileType:
		return &FileStore{filename}
	}
	return nil

}

type FileStore struct {
	Filename string
}

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fs.Filename, fileData, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.Filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, &data)
}
