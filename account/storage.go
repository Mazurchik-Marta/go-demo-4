package account

import (
	"demo/password/crypter"
	"demo/password/output"

	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
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
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type StorageWithDb struct {
	Storage
	db  Db // interface
	enc crypter.Encrypter
}

func NewStorage(db Db, enc crypter.Encrypter) *StorageWithDb {
	file, err := db.Read()
	if err != nil {
		return &StorageWithDb{
			Storage: Storage{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	data := enc.Decrypt(file)
	var storage Storage
	err = json.Unmarshal(data, &storage)
	color.Cyan("Найдено %d аккаунтов", len(storage.Accounts))
	if err != nil {
		output.PrintError("Не удалось преобразовать data.storage")
		return &StorageWithDb{
			Storage: Storage{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &StorageWithDb{
		Storage: storage, 
		db:      db,
		enc:     enc,
	}
}

func (stor *StorageWithDb) AddAccount(acc Account) {
	stor.Accounts = append(stor.Accounts, acc)
	stor.save()
}

func (stor *StorageWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range stor.Accounts {
		isMatched := checker(account, str) 
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
	stor.Accounts = accounts 
	stor.save()
	return isDeleted
}

func (stor *Storage) ToBytes() ([]byte, error) {
	file, err := json.Marshal(stor)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (stor *StorageWithDb) save() {
	stor.UpdatedAt = time.Now()
	data, err := stor.Storage.ToBytes()

	encData := stor.enc.Encrypt(data)
	if err != nil {
		output.PrintError(err)
	}
	stor.db.Write(encData)
}
