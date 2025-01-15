package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {
	fmt.Println(generatePassword(12))

	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAccount := account{
		login:    login,
		password: password,
		url:      url,
	}
	outputPassword(&myAccount)

}

func promtData(promt string) string {
	fmt.Print(promt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(account *account) {
	fmt.Println(account.login, account.password, account.url)
}

func generatePassword(n int) string {
	
	rez := make([]rune, n) // итоговый массив (предслайс рун)
	for i:= range rez {
		rez[i] = letterRunes[rand.IntN(len(letterRunes))] // любой элемент от длинны заготовленных рун ()
	}
	return string(rez)
}
