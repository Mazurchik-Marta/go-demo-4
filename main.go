package main

import (
	"fmt"
	"demo/password/account"
)

func main() {

	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Не верный формат URL или Логин")
		return
	}
	myAccount.OutputPassword()
	fmt.Println(myAccount)
}

func promtData(promt string) string {
	fmt.Print(promt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
