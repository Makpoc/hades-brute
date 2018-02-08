package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ttacon/emoji"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type TZ struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:256"`
	TZ   int
}

const tableName = "tz"

func dbTest(username string) (int, error) {

	return 0, nil
}

// initTable creates the table if it doesn't exist yet
func initTable() error {
	if !db.HasTable(tableName) {
		db.CreateTable(&TZ{})
	}

	if db.Error != nil {
		return fmt.Errorf("failed to create tz table")
	}
	return nil
}

// tzHandler dispatches the command to the appropriate timezone handler
func tzHandler(s *discordgo.Session, m *discordgo.MessageCreate, command commandWithArgs) {
	tzOffset, err := dbTest(m.Author.Username)
	if err != nil {
		fmt.Printf("Failed to do DB request - %v\n", err)
		err = s.MessageReactionAdd(m.ChannelID, m.ID, emoji.Emoji(":thumbsdown:"))
		if err != nil {
			fmt.Printf("Failed to add reaction: %v\n", err)
		}
		return
	}

	_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("TimeZone (GMT): %d", tzOffset))
	if err != nil {
		fmt.Printf("Failed to add reaction: %v\n", err)
	}
}
