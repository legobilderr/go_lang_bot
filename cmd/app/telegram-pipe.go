package app

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"test/internal/telegram"
)

func (a *App) runTelegramPipeline() {
	a.bot.Debug = true
	updates := a.bot.ListenForWebhook("/" + a.bot.Token)

	for {
		select {
		case <-a.ctx.Done():
			a.logger.Info("Finished telegram pipeline")
			return
		case update := <-updates:
			fmt.Println(update.Message.Text)
			switch update.Message.Command() {
			case "good_mornig_radnyli":
				_, err := a.bot.Send(telegram.Nice(update.Message.Chat.ID, telegram.RandomGreetings(update.Message.From.UserName)))
				if err != nil {
					a.logger.Error("couldn't send message", zap.Error(err))
					break
				}

			case "good_morning_pidarasi":
				_, err := a.bot.Send(telegram.Nice(update.Message.Chat.ID, "ДОБРОЕ УТРО ГЕЁЧКИ!"))
				if err != nil {
					a.logger.Error("couldn't send message", zap.Error(err))
				}

				break
			case "puss_deep":
				telegram.NewPuss(a.bot, update.Message.Chat.ID, update.Message.From.UserName, a.config.TelegramApiKey)
				break

			default:
				reply := "Я НЕ ПОНИМАЮ ЧТО ПРОИСХОДИТ !"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				if _, err := a.bot.Send(msg); err != nil {
					a.logger.Error("default error sending message", zap.Error(err))
				}
				break
			}
		}
	}

}
