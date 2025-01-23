package account

import (
	"demo/password/output"
	"encoding/json"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWrite interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWrite
}

type Storage struct {
	// последняя запись как обновлялся
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type StorageWithDb struct {
	Storage
	db Db // интерфейс
}

func NewStorage(db Db) *StorageWithDb {
	file, err := db.Read()
	if err != nil {
		return &StorageWithDb{
			Storage: Storage{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var storage Storage
	//func json.Unmarshal(data []byte, v any) error
	err = json.Unmarshal(file, &storage) // наполнение через указатель
	if err != nil {
		output.PrintError("Не удалось преобразовать")
		return &StorageWithDb{
			Storage: Storage{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &StorageWithDb{
		Storage: storage, // конструировать из исходных данных
		db:      db,
	}
}

func (stor *StorageWithDb) AddAccount(acc Account) {
	stor.Accounts = append(stor.Accounts, acc)
	stor.save()
}

func (stor *StorageWithDb) FindAccountsByURL(url string) []Account {
	var accounts []Account

	for _, account := range stor.Accounts {
		//Contains(s string, substr string) bool
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (stor *StorageWithDb) DelAccountsByURL(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range stor.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	stor.save()
	return isDeleted
}

func (stor *Storage) ToBytes() ([]byte, error) { // остается только с Storage наче при сохранении весь ДБ
	//json.Marshal(v any) ([]byte, error)
	file, err := json.Marshal(stor)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (stor *StorageWithDb) save() {
	stor.UpdatedAt = time.Now()
	data, err := stor.Storage.ToBytes()
	if err != nil {
		output.PrintError(err)
	}
	// создаем базу
	stor.db.Write(data)
}
