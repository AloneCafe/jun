package dto

import "time"

/*
type StarPost struct {

}

type LikePost struct {
	TagID     []int64   `json:"p_tid" db:"p_tid"`
	TagName   []string  `json:"p_tname" db:"p_tname"`
}

type TagPost struct {
	TagID     []int64   `json:"p_tid" db:"p_tid"`
	TagName   []string  `json:"p_tname" db:"p_tname"`
}

type PrivatePost struct {
	Private        bool      `json:"p_private" db:"p_private"`
}

type CategoryPost struct {
	CategoryID     []int64   `json:"p_cid" db:"p_cid"`
	CategoryName   []string  `json:"p_cname" db:"p_cname"`
}

type BanPost struct {
	Banned         bool      `json:"p_banned" db:"p_banned"`
}
*/

type Post struct {
	PID        int64      `json:"p_id" db:"p_id"`
	Title      *string    `json:"p_title" db:"p_title"`
	Desc       *string    `json:"p_desc" db:"p_desc"`
	Body       *string    `json:"p_body" db:"p_body"`
	AuthorID   int64      `json:"u_id" db:"u_id"`
	Keywords   *string    `json:"p_keywords" db:"p_keywords"`
	Type       *string    `json:"p_type" db:"p_type"`
	CreateTime *time.Time `json:"p_create_time" db:"p_create_time"`
	UpdateTime *time.Time `json:"p_update_time" db:"p_update_time"`
	Thumbnails *string    `json:"p_thumbnails" db:"p_thumbnails"`
}

type PostDetail struct {
	Post
	LikeCount    int64 `json:"p_like_cnt" db:"p_like_cnt"`
	UnlikeCount  int64 `json:"p_unlike_cnt" db:"p_unlike_cnt"`
	CommentCount int64 `json:"p_comment_cnt" db:"p_comment_cnt"`
	StarCount    int64 `json:"p_star_cnt" db:"p_star_cnt"`
}
