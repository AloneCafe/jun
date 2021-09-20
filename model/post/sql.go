package post

import (
	"jun/dao"
	"jun/dto"
	"strings"
)

func getAll() ([]dto.Post, error) {
	var posts []dto.Post
	err := dao.QueryN(&posts, "select * from post")
	return posts, err
}

func getAllByLateUpdateTime() ([]dto.Post, error) {
	var posts []dto.Post
	err := dao.QueryN(&posts, "select * from post order by p_update_time desc")
	return posts, err
}

func getAllByLateCreateTime() ([]dto.Post, error) {
	var posts []dto.Post
	err := dao.QueryN(&posts, "select * from post order by p_create_time desc")
	return posts, err
}

func getAllByLatestCreateTime() ([]dto.Post, error) {
	var posts []dto.Post
	sql := `
select * from post a inner join
(select post.p_id, if(
    (exists (select 1 from comment, post where post.p_id = comment.to_p_id)),
    (select comment.c_create_time from comment, post where post.p_id = comment.to_p_id),
    (select post.p_update_time from post)
) as ut from post order by ut desc limit 0,10) b on a.p_id = b.p_id
	`
	err := dao.QueryN(&posts, sql)
	return posts, err
}

// TODO
func getAllBy() ([]dto.Post, error) {
	var posts []dto.Post
	err := dao.QueryN(&posts,
		"select * from post, comment where comment.to_p_id = p_id order by comment.c_update_time desc")
	return posts, err
}

func countAll() (int64, error) {
	var cnt int64
	err := dao.Query1(&cnt, "select count(p_id) from post")
	return cnt, err
}

func getAllByAuthorID(uid int64) ([]dto.Post, error) {
	var posts []dto.Post
	err := dao.QueryN(&posts, "select * from post where u_id = ?", uid)
	return posts, err
}

func searchAllByTitleAndDesc(search string) ([]dto.Post, error) {
	var posts []dto.Post
	var sql string
	for i, word := range strings.Fields(search) {
		if i == 0 {
			sql = `set @v1 := concat('%', 'u', '%'); select * from post where 
			(p_title like @v1 or p_desc like @v1)`
		} else {
			sql += ` and (p_title like @v1 or p_desc like @v1)`
		}
	}

	return posts, err
}
