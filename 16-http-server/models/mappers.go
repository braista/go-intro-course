package models

import "github.com/braista/go-intro-course/16-http-server/internal/database"

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}

func DatabaseUsersToUsers(dbUsers []database.User) []User {
	users := make([]User, len(dbUsers))
	for i, dbUser := range dbUsers {
		users[i] = DatabaseUserToUser(dbUser)
	}
	return users
}

func DatabaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func DatabaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		feeds[i] = DatabaseFeedToFeed(dbFeed)
	}
	return feeds
}
