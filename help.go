package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// helpHandler responds to help command with usage info
func helpHandler(s *discordgo.Session, m *discordgo.MessageCreate, command commandWithArgs) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Some help info comming up")
	if err != nil {
		fmt.Println(err)
	}
}
