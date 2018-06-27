package persistence

import (
	"os"
	"path/filepath"
	"encoding/gob"
)

type persistence struct {
	dataDir string
}

func NewPersistence(dataDir string) *persistence {
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		panic(err)
	}
	return &persistence{
		dataDir: dataDir,
	}
}

func (p *persistence)Save(key string, object interface{}) error {
	file, err := os.Create(filepath.Join(p.dataDir, key))
	defer file.Close()
	if err != nil {
		return err
	}
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(object)
	return err
}

func (p *persistence)Load(key string, object interface{}) error {
	file, err := os.Open(filepath.Join(p.dataDir, key))
	defer file.Close()
	if err != nil {
		return err
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(object)
	return err
}