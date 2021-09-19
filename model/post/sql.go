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

func getAllByUpdateTimeDESC() ([]dto.Post, error) {
	var posts []dto.Post
	err := dao.QueryN(&posts, "select * from post order by p_update_time desc")
	return posts, err
}

func getAllByCreateTimeDESC() ([]dto.Post, error) {
	var posts []dto.Post
	err := dao.QueryN(&posts, "select * from post order by p_create_time desc")
	return posts, err
}

// TODO
func getAllByLastCommentTimeDESC() ([]dto.Post, error) {
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


	err := dao.QueryN(&posts, )
	, word, word)
	return posts, err
}