package modules

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/rain1598/disbash/utils"
)

func Neofetch(flag *[]string, s *discordgo.Session, m *discordgo.MessageCreate) {
	guild, err := s.Guild(m.GuildID)
	if err != nil {
		fmt.Print(err)
	}
	name := "\n[" + m.Author.Username + "@" + guild.Name + "]\n"
	var info string
	for i := 0; i < len(name)+2; i++ {
		info += "-"
	}
	chans, err := s.GuildChannels(m.GuildID)
	if err != nil {
		fmt.Print(err)
	}
	members, err := s.GuildMembers(m.GuildID, "", 1000)
	if err != nil {
		fmt.Print(err)
	}
	membercount := strconv.Itoa(len(members))
	if membercount == "1000" {
		membercount = ">=1000"
	}
	user, err := s.User(guild.OwnerID)
	if err != nil {
		return
	}
	info += "\nOwner: " + user.Username + "#" + user.Discriminator +
		"\nChannels: " + strconv.Itoa(len(chans)) +
		"\nMembers: " + membercount +
		"\nEmojis: " + strconv.Itoa(len(guild.Emojis)) +
		"\nRoles: " + strconv.Itoa(len(guild.Roles)) +
		"\nLocale: " + guild.PreferredLocale +
		"\nRegion: " + guild.Region +
		"\nPremium tier: " + strconv.Itoa(int(guild.PremiumTier)) +
		"\nBoosters: " + strconv.Itoa(guild.PremiumSubscriptionCount) +
		"\nAFK timeout: " + strconv.Itoa(guild.AfkTimeout) + "s" +
		"\nContent filter level: " + strconv.Itoa(int(guild.ExplicitContentFilter))
	img, _ := s.GuildIcon(m.GuildID)
	s.ChannelMessageSend(m.ChannelID, "```ini\n"+utils.GetAscii(32, &img)+name+info+"```")
}