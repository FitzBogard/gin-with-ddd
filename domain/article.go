package domain

type Article struct {
	ID          int64
	Title       string
	UserID      int64
	PublishAt   int64
	LastReplyAt int64
}
