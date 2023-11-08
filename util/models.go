package util

import (
	"time"

	"github.com/google/uuid"
	"github.com/wallacez11/go-rssaggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func DatabaseConvertUser(dbUser database.User) User {

	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdateAt:  dbUser.CreatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

func DatabaseConvertFeed(dbFeed database.Feed) Feed {

	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdateAt:  dbFeed.CreatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}

}

func DatabaseConvertFeedFollow(dbFeed database.FeedFollow) FeedFollows {

	return FeedFollows{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdateAt:  dbFeed.CreatedAt,
		UserID:    dbFeed.UserID,
		FeedId:    dbFeed.FeedID,
	}

}

func DatabaseMultipleFeeds(dbFeed []database.Feed) []Feed {

	feeds := []Feed{}

	for _, dbFeed := range dbFeed {
		feeds = append(feeds, DatabaseConvertFeed(dbFeed))
	}

	return feeds

}

func DatabaseMultipleFeedsFollow(dbFeedFollows []database.FeedFollow) []FeedFollows {

	feedsFollows := []FeedFollows{}

	for _, dbFeed := range dbFeedFollows {
		feedsFollows = append(feedsFollows, DatabaseConvertFeedFollow(dbFeed))
	}

	return feedsFollows

}

func DataBaseConvertPost(dbPost database.Post) Post {

	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdateAt:    dbPost.PublishedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func DatabaseMultiplePost(dbFeedFollows []database.Post) []Post {

	MultiplePosts := []Post{}

	for _, Post := range dbFeedFollows {
		MultiplePosts = append(MultiplePosts, DataBaseConvertPost(Post))
	}

	return MultiplePosts

}
