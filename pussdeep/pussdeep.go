package pussdeep

import (
	"math/rand"
	"time"
)

func Random_deep_pusse() string {

	deep := []string{
		"2 см %s облажает пусей муравья ",
		"10 см %s ну чтож один раз не педафил ",
		"19 см парни с %s можно с разгона влетать ",
		"1 метр если пойдет дождь будем прятаться в %s",
	}

	rand.Seed(time.Now().UnixNano())

	return deep[rand.Intn(len(deep))]
}
