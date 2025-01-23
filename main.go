package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {

	fmt.Println("__Менеджер паролей__")
	// func account.NewStorage(db *files.JsonDB) *account.StorageWithDb
	storage := account.NewStorage(files.NewJsonDB("data.Json"))
	//storage := account.NewStorage(cloud.NewCloudDB("kloi")) - также можно Интерфейс.!

Menu:
	for {
		variant := promtData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт по URL",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант"})
		switch variant {
		case "1":
			createAccount(storage)
		case "2":
			findAccoun(storage)
		case "3":
			deleteAccount(storage)
		default:
			break Menu
		}
	}
}

func findAccoun(stor *account.StorageWithDb) {
	url := promtData([]string{"Введите URL для поиска"})
	accounts := stor.FindAccountsByURL(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func createAccount(stor *account.StorageWithDb) {
	login := promtData([]string{"Введите логин"})
	password := promtData([]string{"Введите пароль"})
	url := promtData([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Не верный формат URL или Логин")
		return
	}
	stor.AddAccount(*myAccount)
}

func deleteAccount(stor *account.StorageWithDb) {
	url := promtData([]string{"Введите URL для удаления"})
	isDeleted := stor.DelAccountsByURL(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не надено")
	}
}

func promtData[T any](promt []T) string {
	for i, line := range promt {
		if i == len(promt) - 1 {
			fmt.Printf(" %v: ",line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
