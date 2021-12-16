package pussdeep

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"test/model"
	"time"

	"github.com/joho/godotenv"
)

func Random_deep_pusse() []string {

	deep := [][]string{
		{"2 см %s ублажает пусей муравья ", "tiny "},
		{"10 см %s ну чтож один раз не педафил , снимай трусишки", "take_your_pants_of"},
		{"19 см парни в %s можно с разгона влетать", "crashed"},
		{"1 метр если пойдет дождь будем прятаться в %s", "Shaquille_ONeal"},
	}

	rand.Seed(time.Now().UnixNano())

	return deep[rand.Intn(len(deep))]
}

func Serch_gif(gifName string) (string, error) {
	err := godotenv.Load()
	var giphyKey string

	if err != nil {

		giphyKey = os.Getenv("GIPHY_KEY")

	} else {

		myEnv, err := godotenv.Read()
		if err != nil {
			log.Panic(err)
		}
		giphyKey = myEnv["GIPHY_KEY"]

	}

	u := "https://api.giphy.com/v1/gifs/search?api_key=" + giphyKey + "&limit=1&q=" + gifName

	resp, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	var gr model.GiffyResponse
	if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
		fmt.Println(err)
		return "", err
	}
	gifLink := gr.Data[0].Images.Downsized.URL
	return gifLink, err

}

func SendRequestTGapi(telegramkey string, ChatID int64, link string) {
	u := fmt.Sprintf("https://api.telegram.org/bot%s/sendAnimation?chat_id=%d&animation=%s",
		telegramkey,
		ChatID,
		link,
	)
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

}
