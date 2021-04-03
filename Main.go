package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var dg *discordgo.Session

func main() {
	token, err := ioutil.ReadFile("token")
	if err != nil {
		fmt.Print("Please place bot token in a file named \"token\" in the install directory")
		return
	}
	dg, err := discordgo.New("Bot " + string(token))
	if err != nil {
		panic(err)
	}
	dg.Open()
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.AddHandler(listener)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sig

	dg.Close()
}
func listener(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}
	command := strings.Split(m.Content, " ")
	switch command[0] {
	case "neofetch":
		fmt.Print(m.GuildID)
		neofetch(&command, s, m)
	default:
		return
	}
	fmt.Print(command)
}
func neofetch(flag *[]string, s *discordgo.Session, m *discordgo.MessageCreate) {
	img, _ := s.GuildIcon(m.GuildID)
	fmt.Print(getascii(44, &img))
}
