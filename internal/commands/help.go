package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type CmdHelp struct{}

func (c *CmdHelp) Invokes() []string {
	return []string{"help", "h"}
}

func (c *CmdHelp) Description() string {
	return "Returns a list of commands and their descriptions."
}

func (c *CmdHelp) AdminRequired() bool {
	return false
}

func (c *CmdHelp) Exec(ctx *Context) (err error) {
	_, err = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    ctx.Message.Author.Username + "#" + ctx.Message.Author.Discriminator,
			IconURL: ctx.Message.Author.AvatarURL(""),
		},
		Title: "Help Command",
		Color: 0x00ff00,
		// the command shall go in the description
		Description: "This is the help command.\n\n" + "Returns this Embed.",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "ping",
				// return the ping command description
				Value: "Returns the Bot's latency",
			},
			{
				Name: "embed",
				// return the embed command description
				Value: "Returns an Test Embed!",
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    ctx.Message.Author.Username + "#" + ctx.Message.Author.Discriminator,
			IconURL: ctx.Message.Author.AvatarURL(""),
		},
		Timestamp: time.Now().Format(time.RFC3339),
	})
	return
}
