package support

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func GetSupportData(url string) ([]SupportData, error) {
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

	var supportData []SupportData
	err = json.Unmarshal(body, &supportData)
	if err != nil {
		return nil, err
	}

	var filteredSupportData []SupportData
	for _, data := range supportData {
		filteredSupportData = append(filteredSupportData, data)
	}

	return filteredSupportData, nil
}
