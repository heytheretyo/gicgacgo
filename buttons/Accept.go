package buttons

import (
	"fmt"
	shared "gicgacgo/shared"
	"time"

	"github.com/bwmarrin/discordgo"
)

func AcceptDuel(s *discordgo.Session, i *discordgo.InteractionCreate, player1 string, player2 string) {
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

	gameID := fmt.Sprintf("%s_%s", player1, player2)

	shared.Players[player1] = &shared.Player{GameId: gameID, Id: player1}
	shared.Players[player2] = &shared.Player{GameId: gameID, Id: player2}

	shared.Games[gameID] = &shared.Game{
		StartedTimestamp: time.Now(),
		Players:          []shared.Player{*shared.Players[player1], *shared.Players[player2]},
		Turn:             shared.RandomizeTurn(),
		PlayerX:          *shared.Players[player1],
		PlayerY:          *shared.Players[player2],
		Game: shared.Board{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
	}

	shared.DisableAllButtons(s, i)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("duel has started between <@%s> and <@%s>! Use `/place <x> <y>` to play.", player1, player2),
		},
	})

	shared.RenderBoard()
}
