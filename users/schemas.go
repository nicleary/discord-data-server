package users

// DiscordUserID Binding to the URI for the get user by ID request
type DiscordUserID struct {
	UserID string `uri:"id" binding:"required"`
}

type NewDiscordUser struct {
	UserID     string `json:"user_id"`
	DateJoined string `json:"date_joined"`
	IsBot      bool   `json:"is_bot,omitempty"`
}

type DiscordUser struct {
	ID         int    `json:"id"`
	UserID     string `json:"user_id"`
	DateJoined string `json:"date_joined"`
	IsBot      bool   `json:"is_bot"`
}
