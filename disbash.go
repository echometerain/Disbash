package main

import (
	"fmt"
	"io/ioutil" //reading files
	"os"
	"os/signal" //good practice for handling stop signals
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rain1598/disbash/modules"
)

var dg *discordgo.Session
var mods = map[string]struct{}{
	"neofetch": {},
	"pinky":    {},
	"shutdown": {},
}

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
	commands := strings.Split(m.Content, " ")
	if _, ok := mods[commands[0]]; !ok {
		return
	}

	fmap := flags(&commands, s, m)
	fmt.Print(*fmap)
	switch (*fmap)["_"] {
	case "neofetch":
		modules.Neofetch(fmap, s, m)
	case "shutdown":
		os.Exit(0)
	default:
		return
	}
}
func flags(commands *[]string, s *discordgo.Session, m *discordgo.MessageCreate) *map[string]string {
	fmap := map[string]string{}
	if len(*commands) == 0 {
		*commands = strings.Split(m.Content, " ")
	}

	last := "_"
	for _, fl := range *commands {
		if fl[0] == '-' {
			if fl[1] == '-' {
				flp := fl[2 : len(fl)-1]
				fmap[flp] = ""
				last = flp

			} else {
				for _, rfl := range fl[1 : len(fl)-1] {
					sfl := string(rfl)
					fmap[sfl] = ""
					last = sfl
				}
			}
		} else {
			fmap[last] = fl
		}
	}
	return &fmap
}
