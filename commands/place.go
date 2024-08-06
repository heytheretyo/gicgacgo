package command

import (
	shared "gicgacgo/shared"

	"github.com/bwmarrin/discordgo"
)

func Place(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.Member.User

	if shared.Players[user.ID] == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "grrr, you have to be in a game to place a marker!!!",
				Flags:   1 << 6,
			},
		})
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "ping pong ding dong ling long king kong",
		},
	})
}
