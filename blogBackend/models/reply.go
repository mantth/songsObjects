package models

type Reply struct {
	ID           uint64 `json:"id" db:"id"`
	CommentID    uint64 `json:"commentID" db:"comment_id"`
	ReplyContent string `json:"reply" db:"reply_content"`
}
