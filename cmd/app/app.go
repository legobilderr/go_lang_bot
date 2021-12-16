package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"test/internal/config"
	zapLogger "test/internal/pkg/logger"
)

type App struct {
	ctx    context.Context
	config *config.Config
	logger *zap.Logger
	lock   *sync.RWMutex
	wg     *sync.WaitGroup
	mux    *http.ServeMux
	bot    *tgbotapi.BotAPI
}

func NewApp(ctx context.Context, config *config.Config) (a *App, err error) {
	a = &App{
		ctx:    ctx,
		config: config,
		logger: zapLogger.NewLogger(),
		lock:   &sync.RWMutex{},
		wg:     &sync.WaitGroup{},
		mux:    http.NewServeMux(),
	}

	a.muxInit()

	if a.bot, err = tgbotapi.NewBotAPI(config.TelegramApiKey); err != nil {
		a.logger.Error("couldn't connect to bot api", zap.Error(err))
		return nil, err
	}

	return
}

func (a *App) Run(ctx context.Context) error {
	a.wg.Add(1)

	go func(ctx context.Context) {
		a.runTelegramPipeline()
		defer a.wg.Done()
	}(ctx)

	return http.ListenAndServe(fmt.Sprintf(":%s", a.config.Port), a.mux)
}

func (a *App) muxInit() {
	a.mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("works"))
		if err != nil {
			return
		}
	})

	a.mux.HandleFunc("/bot/get-users", func(writer http.ResponseWriter, request *http.Request) {
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
