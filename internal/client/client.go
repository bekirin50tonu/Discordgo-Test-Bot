package client

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sun_bot/internal/client/command/slash/admin"
	"sun_bot/internal/client/command/slash/owner"
	"sun_bot/internal/client/command/slash/user"
	"sun_bot/internal/client/config"
	"sun_bot/pkg/command_manager"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func New() {

	config.New("config.yml")

	client, err := discordgo.New("Bot " + config.Config.Yml.Token)
	if err != nil {
		panic(err)
	}
	client.Identify.Intents = discordgo.IntentsAll

	client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	cmd := command_manager.New()
	cmd.Register(owner.Ping())
	cmd.Register(user.GoogleIt())
	cmd.Register(admin.InfoUser())

	client.AddHandler(cmd.AddHandler)

	err = client.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	cmd.Initialize(client)
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	cmd.DeleteAll(client)
	client.Close()
}
