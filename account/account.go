package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createrdAt time.Time
	updatedAt  time.Time
	Account
}

func (acc Account) OutputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {

	rez := make([]rune, n) // итоговый массив (предслайс рун)
	for i := range rez {
		rez[i] = letterRunes[rand.IntN(len(letterRunes))] // любой элемент от длинны заготовленных рун ()
	}
	acc.password = string(rez) // ничего не возвращаем так как это методи передали *
}
func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	// ParseRequestURI(rawURL string)(*url.URL, error)
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL") // если запись не верна ничего не вернем и оибку
	}

	newAcc := &AccountWithTimeStamp{

		createrdAt: time.Now(),
		updatedAt:  time.Now(),
		Account: Account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
