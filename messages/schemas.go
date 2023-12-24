package messages

import "time"

type DiscordMessage struct {
	Contents string    `json:"contents"`
	SentAt   time.Time `json:"sent_at"`
}
