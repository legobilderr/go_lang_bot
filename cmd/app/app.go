package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"test/internal/config"
	zapLogger "test/internal/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

// App - основная структура (модель данных) приложения
type App struct {
	ctx context.Context

	// config - Конфиг приложения с env переменными
	config *config.Config

	// log - Логгер
	logger *zap.Logger

	// lock & mutex примитивы синхронизации в мультипоточных приложениях
	lock *sync.RWMutex
	wg   *sync.WaitGroup

	// mux -  http Router
	mux *http.ServeMux

	// bot - API для работы с Telegram
	bot *tgbotapi.BotAPI
}

// NewApp - конструктор основной структуры
func NewApp(ctx context.Context, config *config.Config) (a *App, err error) {
	a = &App{
		ctx:    ctx,
		config: config,
		logger: zapLogger.NewLogger(),
		lock:   &sync.RWMutex{},
		wg:     &sync.WaitGroup{},
		mux:    http.NewServeMux(),
	}

	if a.bot, err = tgbotapi.NewBotAPI(config.TelegramApiKey); err != nil {
		a.logger.Error("couldn't connect to bot api", zap.Error(err))
		return nil, err
	}
	a.logger.Info("Connected", zap.String("name", a.bot.Self.UserName))

	return
}

// Run - запускает приложение
func (a *App) Run(ctx context.Context) error {
	a.wg.Add(1)
	a.muxInit()
	updates := a.bot.ListenForWebhook("/" + a.bot.Token)

	go func(ctx context.Context) {
		a.runTelegramPipeline(updates)
		defer a.wg.Done()
	}(ctx)

	return http.ListenAndServe(fmt.Sprintf(":%s", a.config.Port), a.mux)
}

// muxInit - инициализирует роутинг при http запросах
func (a *App) muxInit() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("works"))
		if err != nil {
			return
		}
	})

	http.HandleFunc("/bot/get-users", func(writer http.ResponseWriter, request *http.Request) {
		me, err := a.bot.GetMe()
		if err != nil {
			return
		}

		marshal, err := json.Marshal(me)
		if err != nil {
			return
		}
		_, err = writer.Write(marshal)
		if err != nil {
			return
		}
	})

}
