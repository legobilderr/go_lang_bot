package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"test/greetings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	godotenv "github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	var telegramkey string

	if err != nil {

		telegramkey = os.Getenv("TELEAGRAMBOT_KEY")
		port := os.Getenv("PORT")

		if len(port) == 0 {
			port = "8080"
		}
		if err := http.ListenAndServe(":80", nil); err != nil {
			log.Fatal(err)
		}

	} else {

		myEnv, err := godotenv.Read()
		if err != nil {
			log.Panic(err)
		}
		telegramkey = myEnv["TELEAGRAMBOT_KEY"]

	}

	bot, err := tgbotapi.NewBotAPI(telegramkey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	wh, err := tgbotapi.NewWebhook("https://studibot.herokuapp.com:80")

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60
	// wh, err := tgbotapi.NewWebhook("https://studibot.herokuapp.com/" + bot.Token)
	// _, err = bot.SetWebhook(wh)
	if err != nil {
		panic(err)
	}
	_, err = bot.Request(wh)

	if err != nil {
		panic(err)
	}

	info, err := bot.GetWebhookInfo()

	if err != nil {
		panic(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("failed to set webhook: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/")
	// go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates {

		user_name := update.Message.From.UserName

		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {

			case "good_mornig_radnyli":

				bot.Send(nice(update.Message.Chat.ID, fmt.Sprintf(greetings.RandomFormat(), user_name)))

			case "good_morning_pidarasi":

				reply := "ДОБРОЕ УТРО ГЕЁЧКИ!"

				bot.Send(nice(update.Message.Chat.ID, reply))
				break

			default:
				reply := "ДОБРОЕ УТРО ГЕЁЧКИ!"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

				bot.Send(msg)
			}
		}
	}
}

func nice(id int64, message string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(id, message)
}
