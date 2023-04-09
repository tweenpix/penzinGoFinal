package sms

import (
	"final/lib/validator"
	"io/ioutil"
	"strings"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func ReadSMSData(filePath string) ([]SMSData, error) {
	// Чтение содержимого файла
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Разбиваем содержимое файла на строки
	content := string(bytes)
	lines := strings.Split(content, "\n")

	// Создаём slice для хранения результата
	result := make([]SMSData, 0)

	// Обходим строки файла
	for _, line := range lines {
		// Разбиваем строку на поля
		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			continue
		}

		// Проверяем код страны
		country := fields[0]

		if !validator.Countries[country] {
			continue
		}

		// Проверяем уровень сигнала
		if !validator.BandwidthValidator(fields[1]) {
			continue
		}

		// Проверяем провайдера
		provider := fields[3]
		if !validator.Providers[provider] {
			continue
		}
		// Создаём структуру и добавляем её в результат
		data := SMSData{
			Country:      fields[0],
			Bandwidth:    fields[1],
			ResponseTime: fields[2],
			Provider:     fields[3],
		}
		result = append(result, data)
	}
	return result, nil
}
