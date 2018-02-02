package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// sheetHandler responds to help command with usage info
func sheetHandler(s *discordgo.Session, m *discordgo.MessageCreate, command commandWithArgs) {
	var sheetLink = getEnvPropOrDefault("sheet", "")

	_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Here you go **%s**: %s", m.Author.Username, sheetLink))
	if err != nil {
		fmt.Println(err)
	}
}
