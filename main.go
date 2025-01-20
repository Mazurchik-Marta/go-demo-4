package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {

	fmt.Println("__Менеджер паролей__")

Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccoun()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int 
	fmt.Println("Выберите вариант")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт по URL")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func createAccount () {

	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Не верный формат URL или Логин")
		return
	}
	file, err := myAccount.ToBytes() // преобразовали в слайс байт
	if err!=nil {
		fmt.Println("Не удалось преобразовать  в JSON ")
		return
	}
	files.WritaFile(file, "data.Json") // записали в файл
}

func findAccoun()  {
	
}

func deleteAccount() {
	
}

func promtData(promt string) string {
	fmt.Print(promt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
