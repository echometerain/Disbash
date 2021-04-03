package modules

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/rain1598/disbash/utils"
)

func Neofetch(flag *[]string, s *discordgo.Session, m *discordgo.MessageCreate) {
	//dg.Identify.Intents = discordgo.IntentsAll
	img, _ := s.GuildIcon(m.GuildID)
	s.ChannelMessageSend(m.ChannelID, "```yaml\n"+utils.GetAscii(32, &img)+"```")
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
		"\nContent filter level: " + strconv.Itoa(int(guild.ExplicitContentFilter)) + "```"

	s.ChannelMessageSend(m.ChannelID, info)
}
