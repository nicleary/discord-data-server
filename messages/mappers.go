package messages

import (
	"discord-metrics-server/v2/ent"
	"discord-metrics-server/v2/users"
)

func GetUserMentionsFromMessage(messageObject *ent.Message) []string {
	var mentions []string
	for _, mention := range messageObject.Edges.Mentions {
		mentions = append(mentions, mention.UserID)
	}
	return mentions
}

func MessageToSchema(messageObject *ent.Message) DiscordMessageResponse {
	// Get static content
	var response = DiscordMessageResponse{
		ID:             messageObject.ID,
		User:           users.UserToSchema(messageObject.Edges.Sender),
		Contents:       messageObject.Contents,
		SentAt:         messageObject.SentAt.String(),
		MessageID:      messageObject.MessageID,
		ChannelID:      messageObject.ChannelID,
		CreatedAt:      messageObject.CreatedAt.String(),
		UpdatedAt:      messageObject.UpdatedAt.String(),
		UsersMentioned: GetUserMentionsFromMessage(messageObject),
	}
	// If in reply to is not nil, add it's ID
	if messageObject.Edges.InReplyTo != nil {
		response.InReplyTo = messageObject.Edges.InReplyTo.MessageID
	}
	return response
}
