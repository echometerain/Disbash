package modules

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/rain1598/disbash/utils"
)

func Neofetch(fmap *map[string]string, s *discordgo.Session, m *discordgo.MessageCreate) {
	guild, err := s.Guild(m.GuildID)
	if err != nil {
		fmt.Print(err)
	}
	info := "\n[" + guild.Name + "]\n"
	for i := 0; i < len(guild.Name)+2; i++ {
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
		"\nChannels: " + strconv.Itoa(len(chans))
	info += "\nMembers: " + membercount
	info += "\nEmojis: " + strconv.Itoa(len(guild.Emojis)) +
		"\nRoles: " + strconv.Itoa(len(guild.Roles))
	info += "\nLocale: " + guild.PreferredLocale +
		"\nRegion: " + guild.Region +
		"\nPremium tier: " + strconv.Itoa(int(guild.PremiumTier)) +
		"\nBoosters: " + strconv.Itoa(guild.PremiumSubscriptionCount) +
		"\nAFK timeout: " + strconv.Itoa(guild.AfkTimeout) + "s" +
		"\nContent filter level: " + strconv.Itoa(int(guild.ExplicitContentFilter))
	img, _ := s.GuildIcon(m.GuildID)
	s.ChannelMessageSend(m.ChannelID, "```cs\n"+utils.GetAscii(32, &img)+info+"```")
}
