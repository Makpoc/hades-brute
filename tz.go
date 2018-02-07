package main

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ttacon/emoji"

	_ "github.com/lib/pq"
)

const (
	dbUser = "postgres"
	dbName = "tz"
)

func dbTest() error {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbUser, dbName))

	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping failed", err)
		return err
	}

	fmt.Println("Pong")
	return nil
}

func tzHandler(s *discordgo.Session, m *discordgo.MessageCreate, command commandWithArgs) {
	err := dbTest()
	if err != nil {
		err = s.MessageReactionAdd(m.ChannelID, m.ID, emoji.Emoji(":thumbsdown:"))
		if err != nil {
			fmt.Printf("Failed to add reaction: %v\n", err)
		}
		return
	}

	err = s.MessageReactionAdd(m.ChannelID, m.ID, emoji.Emoji(":thumbsup:"))
	if err != nil {
		fmt.Printf("Failed to add reaction: %v\n", err)
	}
}
