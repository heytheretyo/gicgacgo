package command

import (
	"fmt"
	shared "gicgacgo/shared"

	"github.com/bwmarrin/discordgo"
)

func Duel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//mentions target user -> have to accept -> then game starts
	inviter := i.Member.User
	opponent := i.ApplicationCommandData().Options[0].UserValue(s)

	if inviter.ID == opponent.ID {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "lonely ahh, can't even find someone to fight against u (robot op coming soon :3)",
			},
		})
		return
	}

	if inviter.Bot || opponent.Bot {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "robots are not allowed to duel, let them play captcha instead",
			},
		})
		return
	}

	if shared.Players[inviter.ID] != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "bro, you're literally in a game rn. dont leave ur fren like that",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	if shared.Players[opponent.ID] != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "woaw, the player is already in a game, find someone else bro",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	invitation := &discordgo.MessageSend{
		Content: fmt.Sprintf("<@%s> has invited you to a duel. do you accept?", inviter.ID),
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "yes",
						Style:    discordgo.SuccessButton,
						CustomID: fmt.Sprintf("accept_%s_%s", inviter.ID, opponent.ID),
					},
					discordgo.Button{
						Label:    "no",
						Style:    discordgo.DangerButton,
						CustomID: fmt.Sprintf("decline_%s_%s", inviter.ID, opponent.ID),
					},
				},
			},
		},
	}

	s.ChannelMessageSendComplex(i.ChannelID, invitation)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("invited <@%s> to your duel, please wait tightly (if they dont answer within 15s, invitation will expire)", opponent.ID),
		},
	})
}
