package command

import (
	"flag"
	"math/rand"
	"strings"
	"time"
)

type Command struct {
	Dir      *string
	Url      *string
	FileName *string
}

var runes = []rune("abcdefghjijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890-_")

func generateRandString() string {

	rand.Seed(time.Now().UnixNano())

	var builder strings.Builder

	for i := 0; i < 16; i++ {
		builder.WriteRune(runes[rand.Int()%len(runes)])
	}

	return builder.String()
}

func (c *Command) Init() {
	dir := flag.String("p", ".", "help message for flag p - path to save file.")
	host := flag.String("h", "", "help message for flag a - host for download.")
	name := flag.String("n", generateRandString(), "help message for flag n - file name")

	flag.Parse()

	c.Dir = dir
	c.FileName = name
	c.Url = host
}
