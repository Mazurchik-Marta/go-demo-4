package files

import (
	"demo/password/output"
	"os"
	"github.com/fatih/color"
)

type JsonDB struct {
	filename string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		filename: name, 
	}
}

func (db *JsonDB) Read() ([]byte, error) { 
	date, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return date, err
}

func (db *JsonDB) Write(content []byte) { 
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err)
		return
	}
	color.Green("Запись проведена успешно!")
}
