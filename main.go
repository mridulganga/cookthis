package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	channel_ids = map[string]string{}
	mainSession *discordgo.Session
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token := os.Getenv("DISCORD_TOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	go dishScheduler()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "hello":
		s.ChannelMessageSend(m.ChannelID, "Hi Good Person")
	case "cookthis here":
		s.ChannelMessageSend(m.ChannelID, "I will send dishes here")
		channel_ids[m.ChannelID] = ""
		mainSession = s
	case "cookthis stop":
		s.ChannelMessageSend(m.ChannelID, "oh ok jeez, I will stop")
		delete(channel_ids, m.ChannelID)
	}
}

func dishScheduler() {
	for {
		time.Sleep(time.Second * 3)
		for k := range channel_ids {
			mainSession.ChannelMessageSend(k, "cookthis dish")
		}
	}
}
