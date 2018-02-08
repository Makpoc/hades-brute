package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jinzhu/gorm"
	"github.com/ttacon/emoji"
)

const token string = "NDA4NTkzNjk2OTk3NjM4MTQ0.DVSY9Q.IQBtk09A2PWN8HB0AkxGfnbL960"

var botID string
var botPrefix = "."

var (
	backendSecret string
	backendURL    string

	dbUser string
	dbPass string
	dbName string

	db *gorm.DB
)

type CommandHandler interface {
	Register()
	Handle(*discordgo.Session, *discordgo.MessageCreate, commandWithArgs)
	Help() string
}

// commandWithArgs is the message content split into command (the first word) and arguments (everything after that)
type commandWithArgs struct {
	command string
	args    []string
}

// commandHandler is the interface, defining a handler for this bot's commands
type commandHandler interface {
	Handle(s *discordgo.Session, m *discordgo.MessageCreate, command commandWithArgs)
}

// supportedCommands is the map of supported command triggers and their corresponding handlers
var supportedCommands = map[string]func(*discordgo.Session, *discordgo.MessageCreate, commandWithArgs){
	"map":    mapHandler,
	"help":   helpHandler,
	"sheet":  sheetHandler,
	"coffee": coffeeHandler,
	"frosty": frostyHandler,
	"tz":     tzHandler,
}

func main() {
	initEnv()

	// var err error
	// db, err = gorm.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", dbUser, dbName, dbPass))
	// if err != nil {
	// 	printAndExit(err)
	// }
	// defer db.Close()

	dg, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		printAndExit(err)

	}

	u, err := dg.User("@me")
	if err != nil {
		printAndExit(err)
	}

	botID = u.ID

	dg.AddHandler(commandDispatchHandler)

	err = dg.Open()
	if err != nil {
		printAndExit(err)
	}

	fmt.Println("Bot is running")
	<-make(chan struct{})

	return
}

func printAndExit(err error) {
	fmt.Printf("%v\n", err)
	os.Exit(1)
}

// initEnv initializes the application from the environment
func initEnv() {
	backendSecret = getEnvPropOrDefault("secret", "")
	backendURL = getEnvPropOrDefault("backendURL", "http://localhost:8080")

	dbPass = getEnvPropOrDefault("dbPass", "")
	dbName = getEnvPropOrDefault("dbName", "")
	dbUser = getEnvPropOrDefault("dbUser", "")

}

// commandDispatchHandler is the Router for discord commands. It checks if the message is intened to be handled by the bot and if so - delegates it to the appropriate handler function.
func commandDispatchHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !shouldBotAnswer(m.Content) {
		return
	}

	command := parseMessage(m.Content)

	if command.command == "" {
		return
	}

	handlerFunc := supportedCommands[strings.ToLower(strings.TrimPrefix(command.command, botPrefix))]
	if handlerFunc == nil {
		return
	}

	handlerFunc(s, m, command)
}

// parseMessage parses the content of the message and returns it as command and args
func parseMessage(content string) commandWithArgs {
	result := commandWithArgs{}

	cont := strings.Split(content, " ")
	if len(cont) == 0 {
		return result // empty
	}

	result.command = cont[0]

	if len(cont) > 1 {
		result.args = cont[1:]
	}

	return result
}

// shouldBotAnswer checks if the message is intended for the bot, by checking if it's prefixed by botPrefix.
func shouldBotAnswer(message string) bool {
	return strings.HasPrefix(message, botPrefix)
}

// addReaction adds reaction using the :reaction: convention.
func addReaction(s *discordgo.Session, cID, mID, reaction string) error {
	return s.MessageReactionAdd(cID, mID, emoji.Emoji(reaction))
}

// contains checks if a set of strings contains given value
func contains(set []string, val string) bool {
	for _, c := range set {
		if val == c {
			return true
		}
	}
	return false
}

// getEnvPropsOrDefault looks for an environment variable for the given key. If such is found - it returns it, otherwise it returns the provided default value
func getEnvPropOrDefault(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}
