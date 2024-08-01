package request

import (
	"fmt"
	"math/rand"
)

func DotabuffUrl(path string) string {
	domens := []string{
		"https://dotabuff.com",
		"https://it.dotabuff.com",
		"https://ka.dotabuff.com",
		"https://de.dotabuff.com",
		"https://fr.dotabuff.com",
	}

	url := domens[rand.Intn(len(domens))] + path
	fmt.Println(url)
	return url
}
