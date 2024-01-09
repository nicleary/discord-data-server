package messages

type NewDiscordMessage struct {
	UserID      int            `json:"user_id"`
	MessageData DiscordMessage `json:"message_data"`
}

type DiscordMessage struct {
	Contents  string `json:"contents"`
	SentAt    string `json:"sent_at"`
	MessageID int    `json:"message_id"`
}
