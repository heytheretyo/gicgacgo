package main

import (
	command "gicgacgo/commands"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping": command.Ping,
	"duel": command.Duel,
}

func handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}

func handleReady(s *discordgo.Session, r *discordgo.Ready) {
	slog.Info("logged in as",
		slog.String("username", s.State.User.Username),
		slog.String("discriminator", s.State.User.Discriminator))
}
