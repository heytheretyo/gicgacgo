package shared

import "github.com/bwmarrin/discordgo"

func DisableAllButtons(s *discordgo.Session, i *discordgo.InteractionCreate) {
	msg, err := s.ChannelMessage(i.ChannelID, i.Message.ID)
	if err != nil {
		return
	}

	var newComponents []discordgo.MessageComponent
	for _, row := range msg.Components {
		actionRow, ok := row.(*discordgo.ActionsRow)
		if !ok {
			continue
		}

		var newRow discordgo.ActionsRow
		for _, component := range actionRow.Components {
			button, ok := component.(*discordgo.Button)
			if !ok {
				continue
			}

			button.Disabled = true

			newRow.Components = append(newRow.Components, button)
		}

		newComponents = append(newComponents, newRow)
	}

	edited := &discordgo.MessageEdit{
		ID:         i.Message.ID,
		Channel:    i.ChannelID,
		Components: &newComponents,
	}

	_, err = s.ChannelMessageEditComplex(edited)
	if err != nil {
		return
	}
}
