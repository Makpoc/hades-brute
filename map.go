package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// mapCommand is the structure of a command, called using [botPrefix]map
type mapCommand struct {
	args    []string
	author  string
	message []string
}

// mapHandler answers calls to map and map [coord|color] message
func mapHandler(s *discordgo.Session, m *discordgo.MessageCreate, command commandWithArgs) {

	mCommand := parseMapCommand(command)
	mCommand.author = m.Author.Username

	url := fmt.Sprintf("%s/map?secret=%s", backendURL, backendSecret)
	if len(mCommand.args) > 0 {
		url = fmt.Sprintf("%s&coords=%s", url, strings.Join(mCommand.args, ","))
	}

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to get map - got %s. Error was: %v\n", resp.Status, err)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":flushed: Failed to get map - %s", resp.Status))
		return
	}
	defer resp.Body.Close()

	err = sendDiscordResponse(s, m, resp, mCommand)
	if err != nil {
		fmt.Println("Something went wrong while sending Discord response", err)
		return
	}

	err = s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		fmt.Println("Failed to delete trigger message", err)
	}
}

// sendDiscordResponse sends the response from the backend to the discord channel it got the trigger from. It will also add a message to the file in that response, containing the author of the trigger and will delete the original message.
func sendDiscordResponse(s *discordgo.Session, m *discordgo.MessageCreate, resp *http.Response, mCommand mapCommand) error {
	respContentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(respContentType, "image/") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":thinking: Suspecious content type: %s!", respContentType))
		return fmt.Errorf("invalid map format")
	}

	var err error
	if len(mCommand.message) > 0 {
		message := fmt.Sprintf("**%s**: %s", mCommand.author, strings.Join(mCommand.message, " "))
		_, err = s.ChannelFileSendWithMessage(m.ChannelID, message, "map.jpeg", resp.Body)
	} else {
		_, err = s.ChannelFileSendWithMessage(m.ChannelID, fmt.Sprintf("**%s** asked for: ", mCommand.author), "map.jpeg", resp.Body)
	}

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, ":flushed: Failed to fullfil your desire")
		return err
	}

	return nil
}

// parseMapCommand parses the given command into a map command struct
func parseMapCommand(command commandWithArgs) mapCommand {
	var mCommand mapCommand
	if len(command.args) > 0 {
		for i, w := range command.args {
			if isValidArgument(w) {
				mCommand.args = append(mCommand.args, w)
			} else {
				mCommand.message = command.args[i:]
				break
			}
		}
	}

	return mCommand
}

// isValidArgument checks if the provided string is a valid argument (coordinate or color)
func isValidArgument(arg string) bool {
	directions := []string{
		"a1", "a2", "a3", "a4",
		"b1", "b2", "b3", "b4", "b5",
		"c1", "c2", "c3", "c4", "c5", "c6",
		"d1", "d2", "d3", "d4", "d5", "d6", "d7",
		"e2", "e3", "e4", "e5", "e6", "e7",
		"f3", "f4", "f5", "f6", "f7",
		"g4", "g5", "g6", "g7",
	}

	colors := []string{
		"green", "orange", "pink", "yellow",
	}

	arg = strings.ToLower(arg)

	return contains(directions, arg) || contains(colors, arg)
}
