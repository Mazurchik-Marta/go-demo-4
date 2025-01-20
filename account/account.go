package account

import (
	"encoding/json"
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	Login    string `json:"login"` 
	Password string `json:"password"` 
	Url      string `json:"url"` 
	CreaterdAt time.Time `json:"createrdAt"` 
	UpdatedAt  time.Time `json:"updatedAt"` 
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.Login)
}

func (acc *Account) ToBytes() ([]byte, error) {
	//json.Marshal(v any) ([]byte, error)
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) generatePassword(n int) {

	rez := make([]rune, n) // итоговый массив (предслайс рун)
	for i := range rez {
		rez[i] = letterRunes[rand.IntN(len(letterRunes))] // любой элемент от длинны заготовленных рун ()
	}
	acc.Password = string(rez) // ничего не возвращаем так как это методи передали *
}
func NewAccount(login, password, urlString string) (*Account, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	// ParseRequestURI(rawURL string)(*url.URL, error)
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL") // если запись не верна ничего не вернем и оибку
	}

	newAcc := &Account{
		Login:    login,
		Password: password,
		Url:      urlString,
		CreaterdAt: time.Now(),
		UpdatedAt:  time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
