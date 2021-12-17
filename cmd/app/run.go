package app

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Run() {
	ctx := context.Background()
	var telegramkey string = env_load()

	app, err := NewApp(ctx, telegramkey)
	if err != nil {
		panic(err)
	}

	app.Run(ctx, telegramkey)
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
func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hi there! I'm DndSpellsBot!"))
}
