package modules

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rain1598/disbash/utils"
)

func Pinky(fmap *map[string]string, s *discordgo.Session, m *discordgo.MessageCreate) {

	dm := false
	channel, _ := s.Channel(m.ChannelID)
	if channel.Type == discordgo.ChannelTypeDM || channel.Type == discordgo.ChannelTypeGroupDM {
		dm = true
	}
	mention := (*fmap)["pinky"]
	id := utils.Getid(&mention)
	if id == "" {
		s.ChannelMessageSend(m.ChannelID, "Invalid mention")
		return
	}
	user, err := s.GuildMember(m.GuildID, id)
	if err != nil {
		fmt.Print(err)
	}
	if user == nil {
		s.ChannelMessageSend(m.ChannelID, "Invalid mention")
		return
	}
	info := "Username: " + user.User.String()

	if !dm {
		if user.Nick != "" {
			info += "\nNickname: " + user.Nick
		}
		joinedat, _ := user.JoinedAt.Parse()
		info += "\nJoined server: " + joinedat.Format(time.Stamp)
		info += "\nRoles: " + strconv.Itoa(len(user.Roles))
		info += "\nMute: " + bts(user.Mute)
		info += "\nDeaf: " + bts(user.Deaf)
		if string(user.PremiumSince) == "" {
			info += "\nBooster since:" + string(user.PremiumSince)
		}
	}
	info += "\nLocale: " + user.User.Locale
	info += "\nBot: " + bts(user.User.Bot)
	info += "\nVerified email: " + bts(user.User.Verified)
	info += "\nMulti-Factor Auth: " + bts(user.User.MFAEnabled)
	s.ChannelMessageSend(m.ChannelID, "```cs\n"+format(&info)+"```")
}
func bts(b bool) string {
	if b {
		return "yes"
	} else {
		return "no"
	}
}
func format(info *string) (rinfo string) {
	infoarr := strings.Split(*info, "\n")
	for i, val := range infoarr {
		rinfo += val
		for j := len(val); j <= 40; j++ {
			rinfo += " "
		}
		if i%2 == 1 {
			rinfo += "\n"
		}
	}
	return
}
