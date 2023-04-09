package voicecall

import (
	"final/lib/validator"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//RU;61;1083;TransparentCalls;0.92;961;59
/*
alpha-2 — код страны;
текущая нагрузка в процентах;
среднее время ответа;
провайдер;
стабильность соединения;
чистота TTFB-связи;
медиана длительности звонка.
*/

type VoicecallData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	BandProvider string
	Stability    float32
	TTFB         int
	VoicePurity  int
	Length       int
}

func ReadVoicecallData(filePath string) ([]VoicecallData, error) {
	// Чтение содержимого файла
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Разбиваем содержимое файла на строки
	content := string(bytes)
	lines := strings.Split(content, "\n")

	// Создаём slice для хранения результата
	result := make([]VoicecallData, 0)

	// Обходим строки файла
	for _, line := range lines {
		// Разбиваем строку на поля
		fields := strings.Split(line, ";")
		if len(fields) != 8 {
			continue
		}

		// Проверяем код страны
		country := fields[0]

		if !validator.Countries[country] {
			continue
		}

		// Проверяем уровень сигнала
		bandwidth := fields[1]
		if !validator.BandwidthValidator(bandwidth) {
			continue
		}

		// Проверяем провайдера
		provider := fields[3]
		if !validator.BandProvider[provider] {
			continue
		}

		// Проверяем стабильность соединения
		stability64, err := strconv.ParseFloat(fields[4], 32)
		if err != nil {
			continue
		}
		stability := float32(stability64)
		if err != nil {
			log.Fatalln(err)
			continue
		}

		// Проверяем ttfb
		ttfb, err := strconv.Atoi(fields[5])
		if err != nil {
			continue
		}
		// Проверяем чистоту
		voicepurity, err := strconv.Atoi(fields[6])
		if err != nil {
			continue
		}
		// Проверяем медиана длительности звонка
		calllength, err := strconv.Atoi(fields[7])
		if err != nil {
			continue
		}

		// Создаём структуру и добавляем её в результат
		data := VoicecallData{
			Country:      country,
			Bandwidth:    bandwidth,
			ResponseTime: fields[2],
			BandProvider: provider,
			Stability:    stability,
			TTFB:         ttfb,
			VoicePurity:  voicepurity,
			Length:       calllength,
		}
		result = append(result, data)
	}
	return result, nil
}
