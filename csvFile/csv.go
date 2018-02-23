package csvFile

import (
	"github.com/mhcvs2/godatastructure/path"
	"os"
	"encoding/csv"
)

type CSV struct {
	path string
	initData []string
}

func NewCSVFile(path string) *CSV {
	return &CSV{path:path}
}

func (c *CSV) Init(initData ...string) error {
	if exist, err := path.Exists(c.path); err != nil {
		return err
	} else if exist {
		return nil
	} else if f, err := os.Create(c.path); err != nil {
		return err
	} else {
		defer f.Close()
		f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
		w := csv.NewWriter(f)//创建一个新的写入文件流
		defer w.Flush()
		if len(initData) > 0 {
			w.Write(initData)
			c.initData = initData
		} else if len(c.initData) > 0 {
			w.Write(c.initData)
		}
	}
	return nil
}

func (c *CSV) Write(data ...string) error {
	if err := c.Init(); err !=nil {
		return err
	}
	if f, err := os.Open(c.path); err !=nil {
		return err
	} else {
		defer f.Close()
		f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
		w := csv.NewWriter(f)//创建一个新的写入文件流
		defer w.Flush()
		w.Write(data)
	}
	return nil
}

func (c *CSV) WriteAll(data [][]string) error {
	if err := c.Init(); err !=nil {
		return err
	}
	if f, err := os.Open(c.path); err !=nil {
		return err
	} else {
		defer f.Close()
		f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
		w := csv.NewWriter(f)//创建一个新的写入文件流
		defer w.Flush()
		w.WriteAll(data)
	}
	return nil
}
