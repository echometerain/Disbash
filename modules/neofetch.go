package modules

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/rain1598/disbash/utils"
)

func Neofetch(fmap *map[string]string, s *discordgo.Session, m *discordgo.MessageCreate) {
	_, infoOnly := (*fmap)["i"]
	_, fast := (*fmap)["f"]
	_, vorbose := (*fmap)["v"]

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		fmt.Print(err)
	}
	img, err := s.GuildIcon(m.GuildID)
	if err != nil {
		fmt.Print(err)
	}
	info := "```cs\n"
	if !infoOnly {
		info += utils.GetAscii(32, &img)
	}
	info += "\n[" + guild.Name + "]\n"
	for i := 0; i < len(guild.Name)+2; i++ {
		info += "-"
	}
	var chans []*discordgo.Channel
	membercount := ""
	var members []*discordgo.Member
	if !fast {
		chans, err = s.GuildChannels(m.GuildID)
		if err != nil {
			fmt.Print(err)
		}
		members, err = s.GuildMembers(m.GuildID, "", 1000)
		if err != nil {
			fmt.Print(err)
		}
		membercount := strconv.Itoa(len(members))
		if membercount == "1000" {
			membercount = ">=1000"
		}
	}
	user, err := s.User(guild.OwnerID)
	if err != nil {
		return
	}
	info += "\nOwner: " + user.Username + "#" + user.Discriminator
	if !fast {
		info += "\nChannels: "
		if vorbose {
			text := 0
			voice := 0
			for _, val := range chans {
				if (*val).Type == discordgo.ChannelType(2) {
					voice++
				} else {
					text++
				}
			}
			info += strconv.Itoa(voice) + " (voice) " + strconv.Itoa(text) + " (text)"
		} else {
			info += strconv.Itoa(len(chans))
		}
		info += "\nMembers: "
		if vorbose {
			mods := 0
			mems := 0
			for _, val := range members {
				p, err := s.UserChannelPermissions((*val).User.ID, m.ChannelID)
				if err != nil {
					return
				}
				if p == discordgo.PermissionManageMessages {
					mods++
				} else {
					mems++
				}
			}
			info += strconv.Itoa(mods) + " (mods) " + strconv.Itoa(mems) + " (members)"
		} else {
			info += membercount
		}
	}
	info += "\nEmojis: " + strconv.Itoa(len(guild.Emojis)) +
		"\nRoles: " + strconv.Itoa(len(guild.Roles))
	if vorbose {
		info += "\nLocale: " + guild.PreferredLocale
	}
	info += "\nRegion: " + guild.Region +
		"\nPremium tier: " + strconv.Itoa(int(guild.PremiumTier)) +
		"\nBoosters: " + strconv.Itoa(guild.PremiumSubscriptionCount) +
		"\nAFK timeout: " + strconv.Itoa(guild.AfkTimeout) + "s" +
		"\nContent filter level: " + strconv.Itoa(int(guild.ExplicitContentFilter))
	s.ChannelMessageSend(m.ChannelID, utils.GetAscii(32, &img)+info+"```")
}
