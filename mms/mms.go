package mms

import (
	"encoding/json"
	"final/lib/validator"
	"io/ioutil"
	"log"
	"net/http"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func GetMMSData(url string) ([]MMSData, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var mmsData []MMSData
	err = json.Unmarshal(body, &mmsData)
	if err != nil {
		return nil, err
	}

	var filteredMMSData []MMSData
	for _, data := range mmsData {
		if !validator.Providers[data.Provider] && !validator.Countries[data.Country] {
			continue
		}
		filteredMMSData = append(filteredMMSData, data)
	}

	return filteredMMSData, nil
}
