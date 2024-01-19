package messages

import (
	"context"
	"discord-metrics-server/v2/db"
	"discord-metrics-server/v2/ent"
	"discord-metrics-server/v2/ent/message"
)

func GetMessageByID(messageID string) (*ent.Message, *string) {
	client := db.GetClient()

	messageObject, err := client.Message.
		Query().
		Where(message.MessageID(messageID)).
		WithMentions().
		WithInReplyTo().
		WithSender().
		Only(context.Background())

	if err != nil {
		ret := "No message with ID exists"
		return nil, &ret
	}

	return messageObject, nil
}
