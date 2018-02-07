package main

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ttacon/emoji"

	_ "github.com/lib/pq"
)

const tableName = "tz"

func dbTest(username string) (int, error) {
	dbUser := getEnvPropOrDefault("dbUser", "")
	dbPass := getEnvPropOrDefault("dbPass", "")
	dbName := getEnvPropOrDefault("dbName", "")

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPass, dbName))

	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlStatement := `SELECT timezone FROM tz WHERE username=$1;`
	var tz int
	row := db.QueryRow(sqlStatement, username)
	err = row.Scan(&tz)
	if err != nil {
		return 0, err
	}
	return tz, nil
}

func tzHandler(s *discordgo.Session, m *discordgo.MessageCreate, command commandWithArgs) {
	tzOffset, err := dbTest(m.Author.Username)
	if err != nil {
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
