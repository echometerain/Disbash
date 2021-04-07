package modules

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rain1598/disbash/utils"
)

func Neofetch(fmap *map[string]string, s *discordgo.Session, m *discordgo.MessageCreate) {
	_, infoOnly := (*fmap)["i"]
	_, fast := (*fmap)["f"]
	_, vorbose := (*fmap)["v"]
	_, side := (*fmap)["s"]

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		fmt.Print(err)
	}
	img, err := s.GuildIcon(m.GuildID)
	if err != nil {
		fmt.Print(err)
	}
	imgtxt := ""
	if !infoOnly {
		imgtxt = utils.GetAscii(32, &img) + "\n"
	}
	info := "[" + guild.Name + "]\n"
	for i := 0; i < len(guild.Name)+2; i++ {
		info += "-"
	}
	var chans []*discordgo.Channel
	var membercount string
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
		membercount = strconv.Itoa(len(members))
		if membercount == "1000" {
			membercount = ">=1000"
		}
	}
	user, err := s.User(guild.OwnerID)
	if err != nil {
		return
	}
	info += "\nOwner: " + user.String()
	if !fast {
		info += "\nChannels: "
		if vorbose {
			text := 0
			voice := 0
			cats := 0
			for _, val := range chans {
				if (*val).Type == discordgo.ChannelTypeGuildVoice {
					voice++
				} else if (*val).Type == discordgo.ChannelTypeGuildCategory {
					cats++
				} else {
					text++
				}
			}
			info += strconv.Itoa(voice) + " (voice) " + strconv.Itoa(text) + " (text)" +
				"\nCategories: " + strconv.Itoa(cats)
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
				if p&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages ||
					p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator {
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
		"\nBoosters: " + strconv.Itoa(guild.PremiumSubscriptionCount)
	if vorbose {
		info += "\nAFK timeout: " + strconv.Itoa(guild.AfkTimeout) + "s" +
			"\nContent filter level: " + strconv.Itoa(int(guild.ExplicitContentFilter))
	}
	if !infoOnly && side {
		imgarr := strings.Split(imgtxt, "\n")
		infoarr := strings.Split(info, "\n")
		infolen := len(infoarr)
		result := ""
		for i := 0; i < 22; i++ {
			if i < infolen {
				result += imgarr[i] + "   " + infoarr[i] + "\n"
			} else {
				result += imgarr[i] + "\n"
			}
		}
		s.ChannelMessageSend(m.ChannelID, "```yaml\n"+result+"```")
	} else {
		s.ChannelMessageSend(m.ChannelID, "```yaml\n"+imgtxt+info+"```")
	}
}
