package app

import (
	"context"
	"log"
	"net/http"
	"sync"
	"test/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// App - основная структура (модель данных) приложения
type App struct {
	ctx context.Context
	// config - Конфиг приложения с env переменными
	config *config.Config
	// lock & mutex примитивы синхронизации в мультипоточных приложениях
	lock *sync.RWMutex
	wg   *sync.WaitGroup

	// bot - API для работы с Telegram
	bot *tgbotapi.BotAPI
}

// NewApp - конструктор основной структуры
func NewApp(ctx context.Context, config *config.Config) (a *App, err error) {
	a = &App{
		ctx:    ctx,
		config: config,
		lock:   &sync.RWMutex{},
		wg:     &sync.WaitGroup{},
	}

	a.env_load()

	if a.bot, err = tgbotapi.NewBotAPI(config.TelegramApiKey); err != nil {
		log.Printf("couldn't connect to bot api")
		return nil, err
	}
	a.bot.Debug = true
	log.Printf("Authorized on account %s", a.bot.Self.UserName)

	return
}

// Run - запускает приложение
func (a *App) Run(ctx context.Context) {
	updates := a.bot.ListenForWebhook("/" + a.bot.Token)
	a.runTelegramPipeline(updates, a.bot, a.config.TelegramApiKey)
}

func (a *App) env_load() {
	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(":"+a.config.Port, nil)

}
func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hi there! I'm DndSpellsBot!"))
}

// muxInit - инициализирует роутинг при http запросах
// func (a *App) muxInit() {
// 	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
// 		_, err := writer.Write([]byte("works"))
// 		if err != nil {
// 			return
// 		}
// 	})

// 	http.HandleFunc("/bot/get-users", func(writer http.ResponseWriter, request *http.Request) {
// 		me, err := a.bot.GetMe()
// 		if err != nil {
// 			return
// 		}

// 		marshal, err := json.Marshal(me)
// 		if err != nil {
// 			return
// 		}
// 		_, err = writer.Write(marshal)
// 		if err != nil {
// 			return
// 		}
// 	})

// }
