package persistence

import (
	"os"
	"path/filepath"
	"encoding/gob"
	"github.com/mhcvs2/godatastructure/util"
	"sync"
)

type persistence struct {
	dataDir string
	lock sync.RWMutex
}

func NewPersistence(dataDir string) *persistence {
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		panic(err)
	}
	return &persistence{
		dataDir: dataDir,
	}
}

func (p *persistence) Exist(key string) bool {
	if exist, err := util.Exists(filepath.Join(p.dataDir, key)); exist && err != nil {
		return true
	}
	return false
}

func (p *persistence) List() []string {
	res, _ := util.GetSubFiles(p.dataDir)
	keys := make([]string, len(res))
	for i, value := range res {
		keys[i] = filepath.Base(value)
	}
	return keys
}

func (p *persistence)Save(key string, object interface{}) error {
	p.lock.Lock()
	defer p.lock.Unlock()
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
	p.lock.RLock()
	defer p.lock.RUnlock()
	file, err := os.Open(filepath.Join(p.dataDir, key))
	defer file.Close()
	if err != nil {
		return err
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(object)
	return err
}