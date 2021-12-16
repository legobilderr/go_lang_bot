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

func main() {
	var telegramkey string = env_load()

	bot, err := tgbotapi.NewBotAPI(telegramkey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updates := bot.ListenForWebhook("/" + bot.Token)
	komandSwither(updates, bot, telegramkey)
}

func env_load() string {
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
	return telegramkey
}

func komandSwither(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, telegramkey string) {
	message, err := bot.GetWebhookInfo()
	if err == nil {
		fmt.Println(message.IsSet(), message.URL)
	}
	telegramkey = myEnv["TELEAGRAMBOT_KEY"]
	for update := range updates {

		user_name := update.Message.From.UserName

		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {

			case "good_mornig_radnyli":

				bot.Send(pussdeep.Nice(update.Message.Chat.ID, fmt.Sprintf(greetings.RandomFormat(), user_name)))

			case "good_morning_pidarasi":

				reply := "ДОБРОЕ УТРО ГЕЁЧКИ!"

				bot.Send(pussdeep.Nice(update.Message.Chat.ID, reply))
				break

			case "puss_deep":

				pussdeep.NewPuss(bot, update.Message.Chat.ID, user_name, telegramkey)

			default:
				reply := "Я НЕ ПОНИМАЮ ЧТО ПРОИСХОДИТ !"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

				bot.Send(msg)
			}
		}
	}
}
