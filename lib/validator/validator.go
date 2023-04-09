package validator

import (
	"log"
	"strconv"
)

// Создаём map из кодов стран
// Предполагается, что эти коды стран уже известны
var Countries = map[string]bool{
	"RU": true,
	"US": true,
	"GB": true,
	"FR": true,
	"BL": true,
	"AT": true,
	"BG": true,
	"DK": true,
	"CA": true,
	"ES": true,
	"CH": true,
	"TR": true,
	"PE": true,
	"NZ": true,
	"MC": true,
}

// Создаём map из корректных провайдеров
var Providers = map[string]bool{
	"Topolo": true,
	"Rond":   true,
	"Kildy":  true,
}

// Создаём map из корректных провайдеров

var BandProvider = map[string]bool{
	"TransparentCalls": true,
	"E-Voice":          true,
	"JustPhone":        true,
}

var EmailProvider = map[string]bool{
	"Gmail":       true,
	"Yahoo":       true,
	"Hotmail":     true,
	"MSN":         true,
	"Orange":      true,
	"Comcast":     true,
	"AOL":         true,
	"Live":        true,
	"RediffMail":  true,
	"GMX":         true,
	"Proton Mail": true,
	"Yandex":      true,
	"Mail.ru":     true,
}

// проверям уровень сигнала, возвращаем true, false
func BandwidthValidator(field string) bool {

	bandwidth, err := strconv.Atoi(field)

	if err != nil {
		log.Fatalln(`!!! содержится буква в уровне сигнала:` + field)
	}
	if bandwidth < 0 || bandwidth > 100 {
		log.Fatalln(`!!! уровень сигнала вне диапазона: ` + strconv.Itoa(bandwidth))
	}
	if err != nil {
		return false
	}
	return true
}
