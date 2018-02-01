package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ttacon/emoji"
)

const token string = "NDA4NTkzNjk2OTk3NjM4MTQ0.DVSY9Q.IQBtk09A2PWN8HB0AkxGfnbL960"

var botID string
var botPrefix = "."

var (
	backendSecret string
	backendURL    string
)

type commandWithArgs struct {
	command string
	args    []string
}

func main() {
	initEnv()

	dg, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	botID = u.ID

	dg.AddHandler(mapHandler)

	err = dg.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Bot is running")
	<-make(chan struct{})

	return
}

func initEnv() {
	backendSecret = getEnvPropOrDefault("secret", "")
	backendURL = getEnvPropOrDefault("backendURL", "http://159.65.22.117:8080")
}

// mapHandler answers calls to .map and .map [coord] message
func mapHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !shouldBotAnswer(m.Content) {
		return
	}

	command := parseMessage(m.Content)

	if command.command != botPrefix+"map" {
		// not suitable for this handler
		// TODO: Check if there's a way to define this earlier (at handler declaration)
		return
	}

	url := fmt.Sprintf("%s/map?secret=%s", backendURL, backendSecret)

	if len(command.args) > 0 && isValidCoord(command.args[0]) {
		url = fmt.Sprintf("%s&coords=%s", url, command.args[0])
		command.args = command.args[1:]
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get map", err)
		s.ChannelMessageSend(m.ChannelID, ":flushed: Failed to get map")
		return
	}
	defer resp.Body.Close()

	respContentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(respContentType, "image/") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":thinking: Suspecious content type: %s!", respContentType))
		fmt.Println("Invalid map format!")
		return
	}

	if len(command.args) > 0 {
		_, err = s.ChannelFileSendWithMessage(m.ChannelID, strings.Join(command.args[0:], " "), "map.jpeg", resp.Body)
	} else {
		_, err = s.ChannelFileSend(m.ChannelID, "map.jpeg", resp.Body)
	}

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":flushed: Failed to fullfil your desire")
		fmt.Println("Failed to send file!", err)
	}
}

// parseMessage parses the content of the message and returns it as command and args
func parseMessage(content string) commandWithArgs {
	result := commandWithArgs{}

	cont := strings.Split(content, " ")
	if len(cont) < 1 {
		return result // empty
	}

	result.command = cont[0]

	if len(cont) > 1 {
		result.args = cont[1:]
	}

	return result
}

// func dumbShitHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

// 	if m.Author.ID == botID {
// 		return
// 	}

// 	if !shouldBotAnswer(m.Content) {
// 		return
// 	}

// 	err := addReaction(s, m.ChannelID, m.ID, ":beers:")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	if strings.Contains(strings.ToLower(m.Author.Username), "astro") {
// 		addReaction(s, m.ChannelID, m.ID, ":bow:")
// 	}

// 	fmt.Printf("Got message: %s\n", m.Content)
// }

func shouldBotAnswer(message string) bool {
	return strings.HasPrefix(message, botPrefix)
}

// addReaction adds reaction using the :reaction: convention.
func addReaction(s *discordgo.Session, cID, mID, reaction string) error {
	return s.MessageReactionAdd(cID, mID, emoji.Emoji(reaction))
}

// isValidCoord checks if the provided string is a valid coordinate on the map grid
func isValidCoord(coord string) bool {
	directions := []string{
		"a1", "a2", "a3", "a4",
		"b1", "b2", "b3", "b4", "b5",
		"c1", "c2", "c3", "c4", "c5", "c6",
		"d1", "d2", "d3", "d4", "d5", "d6", "d7",
		"e2", "e3", "e4", "e5", "e6", "e7",
		"f3", "f4", "f5", "f6", "f7",
		"g4", "g5", "g6", "g7",
	}

	coord = strings.ToLower(coord)
	for _, c := range directions {
		if coord == c {
			return true
		}
	}
	return false
}

func getEnvPropOrDefault(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}
