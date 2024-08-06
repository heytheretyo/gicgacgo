package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func setupBot() (*discordgo.Session, error) {
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	if botToken == "" {
		log.Fatalf("bot token not provided")
	}

	s, err := discordgo.New("Bot " + botToken)
	if err != nil {
		return nil, err
	}
	return s, nil
}
