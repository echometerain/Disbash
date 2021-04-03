package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
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
	command := strings.Split(m.Content, " ")
	switch command[0] {
	case "neofetch":
		fmt.Print(m.GuildID)
		neofetch(&command, s, m)
	default:
		return
	}
	fmt.Println(command)
}
func neofetch(flag *[]string, s *discordgo.Session, m *discordgo.MessageCreate) {
	dg.Identify.Intents = discordgo.IntentsAll
	img, _ := s.GuildIcon(m.GuildID)
	s.ChannelMessageSend(m.ChannelID, "```yaml\n"+getascii(32, &img)+"```")
	guild, _ := s.Guild(m.GuildID)
	info := "```ini\n[" + guild.Name + "]\n"
	for i := 0; i < len(guild.Name)+2; i++ {
		info += "-"
	}
	chans, _ := s.GuildChannels(m.GuildID)
	members, _ := s.GuildMembers(m.GuildID, "", 1000)
	membercount := strconv.Itoa(len(members))
	if membercount == "1000" {
		membercount = ">=1000"
	}
	user, _ := s.User(guild.OwnerID)
	info += "\nOwner: " + user.Username + "#" + user.Discriminator +
		"\nChannels: " + strconv.Itoa(len(chans)) +
		"\nMembers: " + membercount +
		"\nEmojis: " + strconv.Itoa(len(guild.Emojis)) +
		"\nRoles: " + strconv.Itoa(len(guild.Roles)) +
		"\nLocale: " + guild.PreferredLocale +
		"\nRegion: " + guild.Region +
		"\nPremium tier: " + strconv.Itoa(int(guild.PremiumTier)) +
		"\nBoosters: " + strconv.Itoa(guild.PremiumSubscriptionCount) +
		"\nAFK timeout: " + strconv.Itoa(guild.AfkTimeout) +
		"\nExplicit content filter level: " + strconv.Itoa(int(guild.ExplicitContentFilter)) + "```"

	s.ChannelMessageSend(m.ChannelID, info)

}
