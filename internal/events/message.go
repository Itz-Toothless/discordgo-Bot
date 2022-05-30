package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type MessageHandler struct{}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	channel, err := s.Channel(e.ChannelID)
	if err != nil {
		fmt.Println("Failed getting channel: ", err)
		return
	}
	// Ignore all messages created by the bot itself and messages sent by bots.
	if e.Author.ID == s.State.User.ID || e.Author.Bot {
		return
	}
	fmt.Printf("Message from %s in %s: %s\n", e.Author.String(), channel.Name, e.Content)
}
