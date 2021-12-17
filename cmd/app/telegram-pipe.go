package app

import (
	"fmt"
	"log"
	"test/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *App) runTelegramPipeline(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, telegramkey string) {
	for update := range updates {

		user_name := update.Message.From.UserName

		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {

			case "good_mornig_radnyli":

				bot.Send(telegram.Nice(update.Message.Chat.ID, fmt.Sprintf(telegram.RandomFormat(), user_name)))

			case "good_morning_pidarasi":

				reply := "ДОБРОЕ УТРО ГЕЁЧКИ!"

				bot.Send(telegram.Nice(update.Message.Chat.ID, reply))
				break

			case "puss_deep":

				telegram.NewPuss(bot, update.Message.Chat.ID, user_name, telegramkey)

			default:
				reply := "Я НЕ ПОНИМАЮ ЧТО ПРОИСХОДИТ !"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

				bot.Send(msg)
			}
		}
	}
}
