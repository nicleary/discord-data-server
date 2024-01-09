package messages

import (
	"context"
	"discord-metrics-server/v2/ent"
	"discord-metrics-server/v2/users"
)

func MessageToSchema(messageObject *ent.Message) DiscordMessageResponse {
	// Get user object from message object
	userObject, _ := messageObject.QuerySender().First(context.Background())
	return DiscordMessageResponse{
		ID:        messageObject.ID,
		User:      users.UserToSchema(userObject),
		Contents:  messageObject.Contents,
		SentAt:    messageObject.SentAt.String(),
		MessageID: messageObject.MessageID,
		CreatedAt: messageObject.CreatedAt.String(),
		UpdatedAt: messageObject.UpdatedAt.String(),
	}
}
