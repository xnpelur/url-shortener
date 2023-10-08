package link

import (
	"math"
	"strings"
)

type Link struct {
	Id       int    `json:"id"`
	ShortUrl string `json:"short"`
	Url      string `json:"url"`
}

func (link *Link) Initialize(id int) {
	link.Id = id
	link.ShortUrl = getShortUrl(id)
}

func getShortUrl(id int) string {
	const minLength = 3
	const multiplier = 777

	charset := "bFw9jnNp5HtyYXmE4dQfMxk3KvAeGzr6gU7ZsJh8W2RiC1TcS"
	base := len(charset)

	startId := int(math.Pow(float64(base), minLength-1))
	num := startId + id*multiplier

	var result strings.Builder

	for num > 0 {
		remainder := num % base
		result.WriteString(string(charset[remainder]))
		num /= base
	}

	resultStr := result.String()

	// Reverse string
	runes := []rune(resultStr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
