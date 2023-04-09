package email

import (
	"final/lib/validator"
	"io/ioutil"
	"strconv"
	"strings"
)

//RU;Gmail;120
/*
alpha-2 — код страны;
провайдер;
среднее время доставки писем в миллисекундах.
*/

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

func ReadEmailData(filePath string) ([]EmailData, error) {
	// Чтение содержимого файла
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Разбиваем содержимое файла на строки
	content := string(bytes)
	lines := strings.Split(content, "\n")

	// Создаём slice для хранения результата
	result := make([]EmailData, 0)

	// Обходим строки файла
	for _, line := range lines {
		// Разбиваем строку на поля
		fields := strings.Split(line, ";")
		if len(fields) != 3 {
			continue
		}

		// Проверяем код страны
		country := fields[0]
		if !validator.Countries[country] {
			continue
		}

		// Проверяем провайдера
		provider := fields[1]
		if !validator.EmailProvider[provider] {
			continue
		}
		// Проверяем время доставки
		deliveryTime, err := strconv.Atoi(fields[2])
		if err != nil {
			continue
		}

		// Создаём структуру и добавляем её в результат
		data := EmailData{
			Country:      country,
			Provider:     provider,
			DeliveryTime: deliveryTime,
		}
		result = append(result, data)
	}
	return result, nil
}
