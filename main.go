package main

import (
	"demo/password/account"
	"demo/password/crypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.StorageWithDb){
	"1": createAccount,
	"2": findAccounByURL,
	"3": findAccounByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по имени",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант",
}

// Замыкание
func menuCounter() func() { // возвращаем аннаним фу
	i := 0
	return func() { // ананим фу
		i++
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("__Менеджер паролей__")
	//-----
	// для крипты
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти env. файл")
	}
	//---------------------
	storage := account.NewStorage(files.NewJsonDB("data.storage"), *crypter.NewEncrypter())
	//storage := account.NewStorage(cloud.NewCloudDB("kloi")) - также можно Интерфейс.!
	counter := menuCounter() //cчетчик вызова меню
Menu:
	for {
		counter() // вызов счетчика
		variant := promtData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil { // проверка наличие ключа
			break Menu
		}
		menuFunc(storage)
	}
}

func findAccounByURL(stor *account.StorageWithDb) {
	url := promtData("Введите URL для поиска")
	accounts := stor.FindAccounts(url, func(a account.Account, s string) bool {
		return strings.Contains(a.Url, s) // ананимная функция (которая уже вернула значение)
	})
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func findAccounByLogin(stor *account.StorageWithDb) {
	login := promtData("Введите Логин для поиска")
	accounts := stor.FindAccounts(login, func(a account.Account, s string) bool {
		return strings.Contains(a.Login, s) // ананимная функция (которая уже вернула значение)
	})
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func createAccount(stor *account.StorageWithDb) {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Не верный формат URL или Логин")
		return
	}
	stor.AddAccount(*myAccount)
}

func deleteAccount(stor *account.StorageWithDb) {
	url := promtData("Введите URL для удаления")
	isDeleted := stor.DelAccountsByURL(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не надено")
	}
}

func promtData(promt ...string) string {
	for i, line := range promt {
		if i == len(promt)-1 {
			fmt.Printf(" %v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
