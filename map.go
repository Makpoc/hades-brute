package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type mapCommand struct {
	coords  []string
	message []string
}

// mapHandler answers calls to map and map [coord, ...] message
func mapHandler(s *discordgo.Session, m *discordgo.MessageCreate, command commandWithArgs) {

	mCommand := parseMapCommand(command)

	url := fmt.Sprintf("%s/map?secret=%s", backendURL, backendSecret)
	if len(mCommand.coords) > 0 {
		url = fmt.Sprintf("%s&coords=%s", url, strings.Join(mCommand.coords, ","))
	}

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to get map - got %s. Error was: %v\n", resp.Status, err)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":flushed: Failed to get map - %s", resp.Status))
		return
	}
	defer resp.Body.Close()

	sendDiscordResponse(s, m, resp, mCommand)
}

func sendDiscordResponse(s *discordgo.Session, m *discordgo.MessageCreate, resp *http.Response, mCommand mapCommand) {
	respContentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(respContentType, "image/") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":thinking: Suspecious content type: %s!", respContentType))
		fmt.Println("Invalid map format!")
		return
	}

	var err error
	if len(mCommand.message) > 0 {
		_, err = s.ChannelFileSendWithMessage(m.ChannelID, strings.Join(mCommand.message, " "), "map.jpeg", resp.Body)
	} else {
		_, err = s.ChannelFileSend(m.ChannelID, "map.jpeg", resp.Body)
	}

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":flushed: Failed to fullfil your desire")
		fmt.Println("Failed to send file!", err)
	}
}

// parseMapCommand parses the given command into a map command struct
func parseMapCommand(command commandWithArgs) mapCommand {
	var mCommand mapCommand
	if len(command.args) > 0 {
		for i, w := range command.args {
			if isValidCoord(w) {
				mCommand.coords = append(mCommand.coords, w)
			} else {
				mCommand.message = command.args[i:]
				break
			}
		}
	}

	return mCommand
}
