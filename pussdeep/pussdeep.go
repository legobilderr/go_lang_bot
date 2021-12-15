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

func Random_deep_pusse() string {

	deep := []string{
		"2 см %s ублажает пусей муравья ",
		"10 см %s ну чтож один раз не педафил , снимай трусишки",
		"19 см парни в %s можно с разгона влетать ",
		"1 метр если пойдет дождь будем прятаться в %s",
	}

	rand.Seed(time.Now().UnixNano())

	return deep[rand.Intn(len(deep))]
}

func Serch_gif() (string, error) {
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

	url := "https://api.giphy.com/v1/gifs/search?api_key=" + giphyKey + "&limit=1&q=punch"

	resp, err := http.Get(url)
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

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	gifLink := gr.Data[0].Images.Downsized.URL
	return gifLink, err

}
