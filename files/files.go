package files

import (
	"demo/password/output"
	"fmt"
	"os"

	"github.com/fatih/color"
)
type JsonDB struct {
	filename string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		filename: name, // имя файла
	}
}

func (db *JsonDB) Read() ([]byte, error) { // ает из базы
	//os.ReadFile(name string) ([]byte, error)
	date, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return date, err
}

func (db *JsonDB) Write(content []byte) { // записывает в базу
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
