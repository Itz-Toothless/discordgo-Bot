package commands

import (
	"fmt"
	"time"
)

type CmdPing struct{}

func (c *CmdPing) Invokes() []string {
	return []string{"ping", "p"}
}

func (c *CmdPing) Description() string {
	return "Pong!"
}

func (c *CmdPing) AdminRequired() bool {
	return false
}

func (c *CmdPing) Exec(ctx *Context) (err error) {
	// send a message and edit it after 3 seconds
	msg, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Pinging...")
	if err != nil {
		return err
	}
	// wait for 3 seconds
	<-time.After(3 * time.Second)
	ctx.Session.ChannelMessageEdit(ctx.Message.ChannelID, msg.ID, fmt.Sprintf("Pong! Latency: %.5s ms", ctx.Session.HeartbeatLatency()))
	return
}
