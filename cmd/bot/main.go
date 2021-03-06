package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zekroTutorials/discordgo/internal/commands"
	"github.com/zekroTutorials/discordgo/internal/config"
	"github.com/zekroTutorials/discordgo/internal/events"
)

const defaultConfigName = "../../config/config.json"

var flagConfig = flag.String("c", defaultConfigName, "The location of the config file.")

func main() {
	flag.Parse()

	cfg, err := config.ParseConfigFromJSONFile(*flagConfig)
	if err != nil {
		panic(err)
	}

	s, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		panic(err)
	}

	s.Identify.Intents = discordgo.MakeIntent(
		discordgo.IntentsGuildMembers |
			discordgo.IntentsGuildMessages)

	registerEvents(s)
	registerCommands(s, cfg)

	if err = s.Open(); err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGQUIT)
	<-sc

	// Cleanly close down the Discord session.
	s.Close()
}

func registerEvents(s *discordgo.Session) {
	joinLeaveHandler := events.NewJoinLeaveHandler()
	s.AddHandler(joinLeaveHandler.HandlerJoin)
	s.AddHandler(joinLeaveHandler.HandlerLeave)

	s.AddHandler(events.NewReadyHandler().Handler)
	s.AddHandler(events.NewMessageHandler().Handler)
}

func registerCommands(s *discordgo.Session, cfg *config.Config) {
	cmdHandler := commands.NewCommandHandler(cfg.Prefix)
	cmdHandler.OnError = func(err error, ctx *commands.Context) {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			fmt.Sprintf("Command Execution failed: %s", err.Error()))
	}

	cmdHandler.RegisterCommand(&commands.CmdPing{})
	cmdHandler.RegisterCommand(&commands.CmdHelp{})
	cmdHandler.RegisterCommand(&commands.CmdEmbed{})
	cmdHandler.RegisterMiddleware(&commands.MwPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
