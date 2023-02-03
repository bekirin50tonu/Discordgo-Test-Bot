package user

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sun_bot/pkg/helpers"
	"time"

	"github.com/bwmarrin/discordgo"
)

func GoogleIt() (*discordgo.ApplicationCommand, func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	var (
		defaultMemberPermission int64 = discordgo.PermissionManageServer
		dmPermission                  = false
	)
	app := discordgo.ApplicationCommand{Name: "Wikipedia", Type: discordgo.MessageApplicationCommand, DefaultMemberPermissions: &defaultMemberPermission, DMPermission: &dmPermission}

	callback := func(s *discordgo.Session, in *discordgo.InteractionCreate) {
		search := in.ApplicationCommandData().Resolved.Messages[in.ApplicationCommandData().TargetID].Content
		response, err := http.Get(fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%v", search))
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		var wiki = helpers.Wikipedia{}
		err = discordgo.Unmarshal(body, &wiki)
		if err != nil {
			panic(err)
		}

		footer := fmt.Sprintf("Timestamp: %v", time.Now().Format(time.RFC822))
		embed := discordgo.MessageEmbed{
			Title:       wiki.Title,
			Image:       &discordgo.MessageEmbedImage{URL: wiki.Thumbnail.Source},
			URL:         wiki.ContentUrls.Desktop.Page,
			Description: wiki.Description,
			Color:       0xff0000,
			Footer:      &discordgo.MessageEmbedFooter{Text: footer},
		}
		err = s.InteractionRespond(in.Interaction, &discordgo.InteractionResponse{
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
