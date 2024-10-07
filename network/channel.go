package network

import (
	"log"

	"github.com/clabland/go-homelab-cable/player"
	"github.com/google/uuid"
)

type Channel struct {
	ID   string
	list *player.MediaList
	p    player.Player
}

func NewChannel(list *player.MediaList, name string) *Channel {
	if name == "" {
		name = uuid.New().String()
	}
	return &Channel{
		ID:   name,
		list: list,
	}
}

func (c *Channel) PlayWith(p player.Player) error {
	if c.p != nil {
		if err := c.p.Shutdown(); err != nil {
			return err
		}
	}
	c.p = p

	err := p.Init()
	if err != nil {
		return err
	}
	return p.Play(c.list)
}

func (c *Channel) UpNext() string {
	return c.list.Next()
}

func (c *Channel) Current() string {
	return c.list.Current()
}

func (c *Channel) PlayNext() string {
	_ = c.p.PlayNext()
	return c.Current()
}

func (c *Channel) AdvanceBySeconds(seconds int) int {
	log.Print("Channel.AdvanceBySeconds() called")
	ts := c.p.AdvanceBySeconds(seconds)
	return ts
}
