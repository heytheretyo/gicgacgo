package main

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func registerCommands(s *discordgo.Session) ([]*discordgo.ApplicationCommand, error) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "gives ping pong",
		}, {
			Name: "leaderboard",
			Options: []*discordgo.ApplicationCommandOption{{
				Name:        "type",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices:     []*discordgo.ApplicationCommandOptionChoice{{Name: "local", Value: "local"}, {Name: "global", Value: "global"}},
				Description: "scope of the placement",
				Required:    true,
			}},
			Description: "shows the guild ranking",
		},
		{
			Name:        "duel",
			Description: "invite someone to play against",
			Options: []*discordgo.ApplicationCommandOption{{
				Name:        "username",
				Type:        discordgo.ApplicationCommandOptionUser,
				Description: "person you want to duel",
				Required:    true,
			}},
		},
		{
			Name:        "place",
			Description: "point in the grid to place ur marker",
			Options: []*discordgo.ApplicationCommandOption{{
				Name:        "x",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Description: "row of the tictactoe grid",
				Choices:     []*discordgo.ApplicationCommandOptionChoice{{Name: "one", Value: 1}, {Name: "two", Value: 2}, {Name: "three", Value: 3}},
				Required:    true,
			}, {
				Name:        "y",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Description: "column of the tictactoe grid",
				Choices:     []*discordgo.ApplicationCommandOptionChoice{{Name: "one", Value: 1}, {Name: "two", Value: 2}, {Name: "three", Value: 3}},
				Required:    true,
			},
			},
		},
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "1263867338164539463", v)
		if err != nil {
			slog.Error("cannot create command", slog.String("command", v.Name), slog.Any("error", err))
		}
		registeredCommands[i] = cmd
	}

	return registeredCommands, nil
}

func removeCommands(s *discordgo.Session, commands []*discordgo.ApplicationCommand) {
	for _, v := range commands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "1263867338164539463", v.ID)
		if err != nil {
			slog.Error("cannot delete command", slog.String("command", v.Name), slog.Any("error", err))
		}
	}
}
