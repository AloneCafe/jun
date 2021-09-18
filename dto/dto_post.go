package dto

type PostNative struct {
	ID       int64  `json:"p_id" db:"p_id"`
	Title    string `json:"p_title" db:"p_title"`
	Desc     string `json:"p_desc" db:"p_desc"`
	Body     string `json:"p_body" db:"p_body"`
	AuthorID int64  `json:"u_id" db:"u_id"`
	Keywords string `json:"p_keywords" db:"p_keywords"`
	Type     string `json:"p_type" db:"p_type"`
}
