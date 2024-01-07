package users

import "discord-metrics-server/v2/ent"

func UserToSchema(userObject *ent.User) DiscordUser {
	return DiscordUser{
		ID:         userObject.ID,
		UserID:     userObject.UserID,
		DateJoined: userObject.DateJoined.String(),
		IsBot:      userObject.IsBot,
	}
}
