package messages

import (
	"discord-metrics-server/v2/ent"
	"discord-metrics-server/v2/users"
)

func MessageToSchema(messageObject *ent.Message) DiscordMessageResponse {
	// Get static content
	var response = DiscordMessageResponse{
		ID:        messageObject.ID,
		User:      users.UserToSchema(messageObject.Edges.Sender),
		Contents:  messageObject.Contents,
		SentAt:    messageObject.SentAt.String(),
		MessageID: messageObject.MessageID,
		ChannelID: messageObject.ChannelID,
		CreatedAt: messageObject.CreatedAt.String(),
		UpdatedAt: messageObject.UpdatedAt.String(),
	}
	// If in reply to is not nil, add it's ID
	if messageObject.Edges.InReplyTo != nil {
		response.InReplyTo = messageObject.Edges.InReplyTo.MessageID
	}
	return response
}
