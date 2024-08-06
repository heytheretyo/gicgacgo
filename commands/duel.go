package command

import "github.com/bwmarrin/discordgo"

func Duel(s *discordgo.Session, i *discordgo.InteractionCreate) {

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "ping pong ding dong ling long king kong",
		},
	})
}
