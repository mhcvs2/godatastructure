package persistence

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"github.com/mhcvs2/godatastructure/util"
	"path/filepath"
)

const dataDir  = "/tmp/test/persist"

func TestNewPersistence(t *testing.T) {
	os.Remove(dataDir)
	NewPersistence(dataDir)
	defer os.Remove(dataDir)
	a := assert.New(t)
	exist, err := util.Exists(dataDir)
	a.True(exist)
	a.NoError(err)
}

type TestObject struct {
	aa string
	bb []string
}

func TestPersistence_Save(t *testing.T) {
	p := NewPersistence(dataDir)
	data := &TestObject{
		aa : "mhc",
		bb : []string{"hello", "world"},
	}
	a := assert.New(t)
	key := "test"
	err := p.Save(key, data)
	a.NoError(err)
	exist, err := util.Exists(filepath.Join(dataDir, key))
	a.True(exist)
	a.NoError(err)

	data2 := &TestObject{}
	p.Load(key, data2)
	a.Equal(data2.aa, "mhc")
	a.Equal(len(data2.bb), 2)
}