package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rain1598/disbash/modules"
)

var dg *discordgo.Session

func main() {
	token, err := ioutil.ReadFile("token")
	if err != nil {
		fmt.Print("Please place bot token in a file named \"token\" in the install directory")
		return
	}
	dg, err = discordgo.New("Bot " + string(token))
	if err != nil {
		panic(err)
	}
	dg.Open()
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.AddHandler(listener)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sig

	dg.Close()
}
func listener(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}
	fmap := flags(&m.Content, s, m)
	switch (*fmap)["_"] {
	case "neofetch":
		go modules.Neofetch(fmap, s, m)
	default:
		return
	}
}
func flags(content *string, s *discordgo.Session, m *discordgo.MessageCreate) (fmap *map[string]string) {
	var command []string = strings.Split(m.Content, " ")
	last := "_"
	for _, fl := range command {
		if fl[0] == '-' {
			if fl[1] == '-' {
				flp := fl[2 : len(fl)-1]
				(*fmap)[flp] = ""
				last = flp

			} else {
				for _, rfl := range fl[1 : len(fl)-1] {
					sfl := string(rfl)
					(*fmap)[sfl] = ""
					last = sfl
				}
			}
		} else {
			(*fmap)[last] = fl
		}
	}
	return
}
