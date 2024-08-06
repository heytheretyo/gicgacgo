package command

import (
	"github.com/bwmarrin/discordgo"
)

func Leaderboard(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// TODO: gonna be used after i have a db setted up
	// options := i.ApplicationCommandData().Options

	// should be some database ig
	leaderboard := []struct {
		Name  string
		WL    string
		Games string
	}{
		{"JohnDoe#88", "88%", "213"},
		{"MaryDoe#82", "92%", "221"},
		{"Picasso#444", "100%", "112"},
		{"Jeff#4424", "100%", "444"},
		{"ForniteLuvr#556", "100%", "1241"},
		{"MoneyUp#1001", "144%", "232"},
		{"OnePiece#22", "87%", "231"},
		{"McChicken#222", "99%", "213"},
		{"PhoneMe#996", "88%", "432"},
		{"Waffle#3512", "87%", "321"},
	}

	embed := &discordgo.MessageEmbed{
		Title: "global leaderboard top 10",
		Color: 0x00ff00,
	}

	for _, player := range leaderboard {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   player.Name,
			Value:  "W/L: " + player.WL + " | Games: " + player.Games,
			Inline: false,
		})
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}
