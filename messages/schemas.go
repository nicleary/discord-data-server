package messages

import (
	"discord-metrics-server/v2/users"
	"fmt"
)

type DiscordMessageID struct {
	MessageID string `uri:"id" binding:"required"`
}

type DiscordMessageQuery struct {
	PageSize      int `form:"page_size"`
	PageNumber    int `form:"page_number"`
	MetricsUserID int `form:"metrics_user_id"`
}

func (q DiscordMessageQuery) validate() *string {
	fmt.Print(q.PageNumber)
	if q.PageSize > 100 {
		ret := "Page size cannot be greater than 100"
		return &ret
	}
	if q.PageSize < 1 {
		ret := "Page size must be at least 1"
		return &ret
	}
	if q.PageNumber < 1 {
		ret := "Page number must be at least 1"
		return &ret
	}
	return nil
}

type DiscordMessage struct {
	UserID    string `json:"user_id"`
	Contents  string `json:"contents"`
	SentAt    string `json:"sent_at"`
	MessageID string `json:"message_id"`
}

type DiscordMessageResponse struct {
	ID        int               `json:"id"`
	User      users.DiscordUser `json:"user"`
	Contents  string            `json:"contents"`
	SentAt    string            `json:"sent_at"`
	MessageID string            `json:"message_id"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}
