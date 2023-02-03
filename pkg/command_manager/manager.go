package command_manager

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lus/dgc"
)

type CommandHandler func(*discordgo.Session, *discordgo.InteractionCreate)

type CommandManager struct {
	Commands []Command
	Router   *dgc.Router
}

type Command struct {
	ID       string
	App      *discordgo.ApplicationCommand
	Callback CommandHandler
}

func New() *CommandManager {

	router := dgc.Create(&dgc.Router{
		Prefixes: []string{"!"},
	})

	return &CommandManager{
		Commands: []Command{},
		Router:   router,
	}

}

func (cmd *CommandManager) Register(app *discordgo.ApplicationCommand, callback func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	cmd.Commands = append(cmd.Commands, Command{App: app, Callback: callback, ID: ""})
}

func (cmd *CommandManager) RegisterCmd(command *dgc.Command) {
	cmd.Router.RegisterCmd(command)
}

func (cmd *CommandManager) Initialize(s *discordgo.Session) {
	start := time.Now()
	for id, command := range cmd.Commands {
		rcmd, er := s.ApplicationCommandCreate(s.State.User.ID, "", command.App)
		if er != nil {
			panic(er)
		}
		cmd.Commands[id].ID = rcmd.ID
	}
	cmd_size := len(cmd.Commands)
	stop := time.Now().Sub(start).Seconds()
	title := fmt.Sprintf("Registered Slash Command: %d\tTime: %v sµs", cmd_size, stop)
	fmt.Println(title)
	start = time.Now()
	cmd.Router.Initialize(s)
	cmd_size = len(cmd.Router.Commands)
	stop = time.Now().Sub(start).Seconds()
	title = fmt.Sprintf("Registered Command: %d\tTime: %v sµs", cmd_size, stop)
	fmt.Println(title)
}

func (cmd *CommandManager) DeleteAll(s *discordgo.Session) {
	for _, handledCmd := range cmd.Commands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", handledCmd.ID)
		if err != nil {
			log.Fatalf("Cannot delete slash command %q: %v", handledCmd.App.Name, err)
		}
	}
}

func (cmd *CommandManager) FindByID(Id string) (Command, error) {
	for _, c := range cmd.Commands {
		if c.ID == Id {
			return c, nil
		}
	}
	return Command{}, errors.New("not matched")
}

func (cmd *CommandManager) AddHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	handledCmd, err := cmd.FindByID(i.ApplicationCommandData().ID)
	if err != nil {
		cmd.errorHandler(s, i)
		return
	}

	handledCmd.Callback(s, i)
}

func (manager *CommandManager) errorHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	fmt.Println("Komut Bulunamadı")
}

// https://discord.com/oauth2/authorize?client_id=959919091248943164&permissions=1342532688&scope=bot%20applications.commands&redirect_uri=https%3A%2F%2Fdroplet.gg&response_type=code
