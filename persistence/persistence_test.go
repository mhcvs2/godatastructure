package persistence

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"github.com/mhcvs2/godatastructure/util"
	"path/filepath"
	"fmt"
)

const dataDir  = "/tmp/test/persist"

func TestNewPersistence(t *testing.T) {
	os.RemoveAll(dataDir)
	NewPersistence(dataDir)
	defer os.RemoveAll(dataDir)
	a := assert.New(t)
	exist, err := util.Exists(dataDir)
	a.True(exist)
	a.NoError(err)
}

type TestObject struct {
	Aa string
	Bb []string
}

func TestPersistence_Save(t *testing.T) {
	p := NewPersistence(dataDir)
	defer os.RemoveAll(dataDir)
	data := &TestObject{
		Aa : "mhc",
		Bb : []string{"hello", "world"},
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
	a.Equal(data2.Aa, "mhc")
	a.Equal(len(data2.Bb), 2)
}

func ExamplePersistence_List() {
	p := NewPersistence(dataDir)
	defer os.RemoveAll(dataDir)
	data := &TestObject{
		Aa : "mhc",
		Bb : []string{"hello", "world"},
	}
	p.Save("test1", data)
	p.Save("test2", data)
	fmt.Println(p.List())
	//Output: [test1 test2]
}