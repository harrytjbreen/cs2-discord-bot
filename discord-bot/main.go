package main

import (
	"discord-bot/handlers"
	"discord-bot/util"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type CommandHandler func(args []string) string

var commandHandlers = map[string]CommandHandler{
	"!ping":       handlers.Ping,
	"!report":     handlers.Report,
	"!rankstatus": handlers.RankStatus,
}

func initialise() func() {

	token := util.GetEnvVar("DISCORD_TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return nil
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()

	if err != nil {
		fmt.Println("Error opening connection:", err)
		return nil
	}

	fmt.Println("Bot is now running and listening for messages")

	return func() {
		err := dg.Close()
		if err != nil {
			return
		}
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		fmt.Println("Ignoring message sent by me!")
		return
	}

	// check if content is empty
	if m.Content == "" {
		fmt.Println("Ignoring empty message")
		return
	}

	words := strings.Split(m.Content, " ")

	if handler, ok := commandHandlers[words[0]]; ok {
		response := handler(words)

		if response != "" {
			fmt.Println("Sending message:", response)
			_, err := s.ChannelMessageSend(m.ChannelID, response)
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
		}
	} else {
		fmt.Println("Ignoring message", m.Content)
	}
}

func main() {
	fmt.Println("Hello, World!")
	closeSession := initialise()

	defer closeSession()

	// Keep the program running
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
