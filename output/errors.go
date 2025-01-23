package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	// или ок натация
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки %d", t)
	default:
		color.Red("Неизвестный тип ошибки")
	}

	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
		return
	}
}
