package billing

import (
	"io/ioutil"
)

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

func ReadBillingData(filePath string) ([]BillingData, error) {

	result := make([]BillingData, 0)

	// Чтение содержимого файла
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var mask uint8 = 0
	for i := 0; i < len(data); i++ {
		if data[i] == '1' {
			mask += 1 << (len(data) - i - 1)
		}
	}

	billingData := BillingData{
		CreateCustomer: (mask & (1 << 0)) != 0,
		Purchase:       (mask & (1 << 1)) != 0,
		Payout:         (mask & (1 << 2)) != 0,
		Recurring:      (mask & (1 << 3)) != 0,
		FraudControl:   (mask & (1 << 4)) != 0,
		CheckoutPage:   (mask & (1 << 5)) != 0,
	}
	// Выводим результаты

	result = append(result, billingData)

	// fmt.Println("CreateCustomer:", billingData.CreateCustomer)
	// fmt.Println("Purchase:", billingData.Purchase)
	// fmt.Println("Payout:", billingData.Payout)
	// fmt.Println("Recurring:", billingData.Recurring)
	// fmt.Println("FraudControl:", billingData.FraudControl)
	// fmt.Println("CheckoutPage:", billingData.CheckoutPage)

	return result, nil
}
