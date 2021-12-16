package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"test/greetings"
	"test/pussdeep"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	godotenv "github.com/joho/godotenv"
)

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hi there! I'm DndSpellsBot!"))
}

type app struct {
}

func main() {

	err := godotenv.Load()
	var telegramkey string

	if err != nil {

		telegramkey = os.Getenv("TELEAGRAMBOT_KEY")
		port := os.Getenv("PORT")
		http.HandleFunc("/", MainHandler)
		go http.ListenAndServe(":"+port, nil)

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

	updates := bot.ListenForWebhook("/" + bot.Token)

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

			case "puss_deep":
				var pussiAnswer []string
				pussiAnswer = pussdeep.Random_deep_pusse()
				bot.Send(nice(update.Message.Chat.ID, fmt.Sprintf(pussiAnswer[0], user_name)))
				link, err := pussdeep.Serch_gif(pussiAnswer[1])
				if err != nil {
					log.Panic(err)
				}
				pussdeep.SendRequestTGapi(telegramkey, update.Message.Chat.ID, link)

			default:
				reply := "Я НЕ ПОНИМАЮ ЧТО ПРОИСХОДИТ !"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

				bot.Send(msg)
			}
		}
	}
}

func nice(id int64, message string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(id, message)
}
