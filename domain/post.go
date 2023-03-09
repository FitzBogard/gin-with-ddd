package domain

type Post struct {
	ID          int64
	ArticleID   int64
	UserID      int64
	ReplyPostID int64 // if it's an initial floor, reply post id is 0
	ReplyUserID int64 // if it's an initial floor, reply user id is 0
	Content     string
	Attachment  string // image or other type of files
	PostAt      int64
}
