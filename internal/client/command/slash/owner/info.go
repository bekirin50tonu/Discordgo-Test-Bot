package owner

import (
	"time"

	global_config "sun_bot/internal/client/config"
	"sun_bot/pkg/helpers"

	"github.com/bwmarrin/discordgo"
)

func Ping() (*discordgo.ApplicationCommand, func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	var (
		defaultMemberPermission int64 = discordgo.PermissionManageServer
		dmPermission                  = true
	)
	app := discordgo.ApplicationCommand{Name: "ping", Type: discordgo.ChatApplicationCommand, Description: "Information Command",
		DefaultMemberPermissions: &defaultMemberPermission, DMPermission: &dmPermission}

	callback := func(s *discordgo.Session, in *discordgo.InteractionCreate) {

		// spew.Dump(in)
		if !helpers.Contains(global_config.Config.Yml.Owners, in.Member.User.ID) {
			return
		}

		footer := "System Information"
		embed := discordgo.MessageEmbed{
			Title:     "System Information",
			Timestamp: time.Now().Format(time.RFC3339),
			Color:     0xff0000,
			Footer:    &discordgo.MessageEmbedFooter{Text: footer},
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Operating System", Value: global_config.Runtime.OSInfo.GoOS, Inline: true},
				{Name: "Kernel", Value: global_config.Runtime.OSInfo.Kernel, Inline: true},
				{Name: "CPU Core", Value: global_config.Runtime.OSInfo.Core, Inline: true},
				{Name: "Platform", Value: global_config.Runtime.OSInfo.Platform, Inline: true},
				{Name: "Started At", Value: time.Now().Sub(global_config.Runtime.StartTime).String(), Inline: true},
			},
		}
		err := s.InteractionRespond(in.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds:          []*discordgo.MessageEmbed{&embed},
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
		if err != nil {
			panic(err)
		}
	}

	return &app, callback
}
