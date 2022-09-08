package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

var Token string
var ChatID string
var OrderNumber string

func main() {
	file, err := openLogFile("log.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	log.Println("log file created")

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	Token = os.Getenv("TOKEN")
	ChatID = os.Getenv("CHAT_ID")
	OrderNumber = os.Getenv("ORDER_NUMBER")

	c := cron.New()
	c.AddFunc("* * 1 * *", func() { triggerIkea() })
	c.Start()

	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			triggerIkea()
		}
	}

}

func triggerIkea() {
	data := []Payload{
		{
			OperationName: "DeliveryMethodsOrder",
			Variables: Variables{
				OrderNumber:   OrderNumber,
				ReceiptNumber: "",
				LiteID:        "",
			},
			Query: queryV2,
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://cssom-prod.ingka.com/purchase-history/graphql", body)
	if err != nil {
		log.Println(err.Error())
	}
	req.Header.Set("Accept-Language", "en-in")
	req.Header.Set("Apollographql-Client-Name", "en-IN")
	req.Header.Set("Apollographql-Client-Version", "8.7.10")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		response := []ResponseBodyV2{}
		json.NewDecoder(resp.Body).Decode(&response)
		b, err := json.MarshalIndent(response, "", "  ")

		if err != nil {
			log.Println(err)
		}
		log.Println(string(b))
		statusV1 := response[0].Data.Order.DeliveryMethods[0].Status
		statusV2 := response[0].Data.Order.DeliveryMethods[0].StatusV2

		statusV2JSON, err2 := json.MarshalIndent(statusV2, "", "  ")
		if err2 != nil {
			log.Println(err)
		}
		SendMessage(statusV1)
		SendMessage(string(statusV2JSON))

	}
}

func getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", Token)
}

func SendMessage(text string) (bool, error) {
	var err error
	var response *http.Response

	url := fmt.Sprintf("%s/sendMessage", getUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id": ChatID,
		"text":    text,
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	log.Printf("Message '%s' was sent", text)
	log.Printf("Response JSON: %s", string(body))

	return true, nil
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
