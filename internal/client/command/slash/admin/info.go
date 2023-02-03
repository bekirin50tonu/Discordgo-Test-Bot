package admin

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
)

func InfoUser() (*discordgo.ApplicationCommand, func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	var (
		defaultMemberPermission int64 = 0
		dmPermission                  = false
	)
	app := discordgo.ApplicationCommand{Name: "User Information", Type: discordgo.UserApplicationCommand,
		DefaultMemberPermissions: &defaultMemberPermission, DMPermission: &dmPermission}

	callback := func(s *discordgo.Session, in *discordgo.InteractionCreate) {

		mention := in.ApplicationCommandData().Resolved.Members[in.ApplicationCommandData().TargetID]
		user := in.ApplicationCommandData().Resolved.Users[in.ApplicationCommandData().TargetID]
		spew.Dump(user)

		var roles = ""

		for _, v := range mention.Roles {
			roles = fmt.Sprintf("%v\n<@&%v>", roles, v)
		}

		if roles == "" {
			roles = "None"
		}

		tiers := make(map[int]string, 0)
		tiers[0] = "None"
		tiers[1] = "Nitro Classic"
		tiers[2] = "Nitro"
		tiers[3] = "Nitro Basic"

		embed := discordgo.MessageEmbed{
			Timestamp: time.Now().Format(time.RFC3339),
			Author:    &discordgo.MessageEmbedAuthor{Name: fmt.Sprintf("%v#%s", user.Username, user.Discriminator), IconURL: user.AvatarURL("")},
			Title:     "User Info",
			Color:     user.AccentColor,
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Roles", Value: roles, Inline: true},
				{Name: "Roles", Value: tiers[user.PremiumType], Inline: true},
				{Name: "Is Support Server", Value: "test", Inline: true},
			},
		}
		err := s.InteractionRespond(in.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{&embed},
			},
		})
		if err != nil {
			panic(err)
		}
	}

	return &app, callback
}
