package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type CmdEmbed struct{}

func (c *CmdEmbed) Invokes() []string {
	return []string{"embed", "e"}
}

func (c *CmdEmbed) Description() string {
	return "Returns an Test Embed!"
}

func (c *CmdEmbed) AdminRequired() bool {
	return false
}

func (c *CmdEmbed) Exec(ctx *Context) (err error) {
	// send an embed
	_, err = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    ctx.Message.Author.Username + "#" + ctx.Message.Author.Discriminator,
			IconURL: ctx.Message.Author.AvatarURL(""),
		},
		Title: "Test Embed",
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Field 1",
				Value:  "Value 1",
				Inline: false,
			},
			{
				Name:   "Field 2",
				Value:  "Value 2",
				Inline: true,
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
