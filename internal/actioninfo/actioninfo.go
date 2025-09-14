package actioninfo

import (
	"fmt"
)

// DataParser — интерфейс для тренировки и прогулки
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info — выводит информацию обо всех активностях
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// парсим данные
		if err := dp.Parse(data); err != nil {
			fmt.Printf("Ошибка парсинга: %v\n", err)
			continue
		}

		// формируем строку
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Ошибка формирования информации: %v\n", err)
			continue
		}

		// выводим результат
		fmt.Println(info)
	}
}
