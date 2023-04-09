package main

import (
	"final/billing"
	"final/email"
	"final/mms"
	"final/sms"
	"final/support"
	"final/voicecall"
	"fmt"
	"log"
)

const baseDir = "../skillbox-diploma/"

const smsFile = baseDir + "sms.data"
const voiceFile = baseDir + "voice.data"
const emailFile = baseDir + "email.data"
const billingFile = baseDir + "billing.data"

const serverMMS = "http://127.0.0.1:8383/mms"
const serverSupport = "http://127.0.0.1:8383/support"

func main() {

	data, err := sms.ReadSMSData(smsFile)
	if err != nil {
		log.Fatalf("косяк с %s выходим...", smsFile)
	}
	fmt.Printf("\n\n%s:\n", smsFile)
	fmt.Println(data)

	resMMS, err := mms.GetMMSData(serverMMS)
	if err != nil {
		log.Print("косяк с MMS выходим...")
	} else {
		fmt.Println("\n\nMMS data:")
		fmt.Println(resMMS)
	}

	resVoice, err := voicecall.ReadVoicecallData(voiceFile)
	if err != nil {
		log.Fatalf("косяк с %s выходим...", voiceFile)
	} else {
		fmt.Printf("\n\n%s:\n", voiceFile)
		fmt.Println(resVoice)
	}

	resEmail, err := email.ReadEmailData(emailFile)
	if err != nil {
		log.Fatalf("косяк с %s выходим...", emailFile)
	} else {
		fmt.Printf("\n\n%s:\n", emailFile)
		fmt.Println(resEmail)
	}

	resBilling, err := billing.ReadBillingData(billingFile)
	if err != nil {
		log.Fatalf("косяк с %s выходим...", billingFile)
	} else {
		fmt.Printf("\n\n%s:\n", billingFile)
		fmt.Println(resBilling)
	}

	resSupport, err := support.GetSupportData(serverSupport)
	if err != nil {
		log.Print("косяк с Support выходим...")
	} else {
		fmt.Println("\n\nSupport data:")
		fmt.Println(resSupport)
	}
}
