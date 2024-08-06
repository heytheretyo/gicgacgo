package buttons

import (
	"fmt"
	shared "gicgacgo/shared"

	"github.com/bwmarrin/discordgo"
)

func RejectDuel(s *discordgo.Session, i *discordgo.InteractionCreate, player1 string, player2 string) {
	if i.Member.User.ID != player2 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "respectfully, you are not authorized to respond to this invitation",
				Flags:   1 << 6,
			},
		})
		return
	}

	shared.DisableAllButtons(s, i)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("<@%s> declined the duel invitation.", player2),
		},
	})
}
