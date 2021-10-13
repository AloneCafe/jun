package post

import (
	"github.com/jmoiron/sqlx"
	"jun/dao"
	"jun/dto"
	"jun/utils/dexss"
)

/*
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
	sql :=
	`
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
*/

var (
	sql = `
select
    a.*,
    (select count(like_post.lp_id) from like_post where like_post.p_id = a.p_id and not like_post.lp_neg) as p_like_cnt, # 正赞
    (select count(like_post.lp_id) from like_post where like_post.p_id = a.p_id and like_post.lp_neg) as p_unlike_cnt,    # 负赞
    (select count(comment.c_id) from comment where comment.to_p_id = a.p_id
                                               and if(?, not comment.c_folded, true)
                                               and if(?, not exists (select 1 from ban_comment where ban_comment.c_id = comment.c_id), true)
                                               and if(?, not exists (select 1 from ban_user, comment where comment.u_id = ban_user.u_id), true)
    ) as p_comment_cnt,    # 评论（看传参是否限制判定）
    (select count(star_post.sp_id) from star_post where star_post.to_p_id = a.p_id) as p_star_cnt # 收藏数
from post a inner join (
    select f_post.p_id,
           if(
                   (exists (select 1 from comment, post p where p.p_id = comment.to_p_id)
                       and if(?, not exists (select 1 from ban_comment, comment where comment.to_p_id = f_post.p_id and comment.c_id = ban_comment.c_id), true)
                       and if(?, not exists (select 1 from ban_user, comment where comment.to_p_id = f_post.p_id and comment.u_id = ban_user.u_id), true)
                       and if(?, not (select comment.c_folded from comment where comment.to_p_id = f_post.p_id), true)
                       ), # 考虑评论被封禁、折叠的情况，如果被封禁，就不应该计算其时间
                   (select comment.c_create_time from comment, post p where p.p_id = comment.to_p_id),
                   (select p.p_update_time from post p where p.p_id = f_post.p_id)
               ) as ut from post f_post
    where
        if (? is null or ? = '', f_post.p_title like '%', f_post.p_title regexp ?)
      and
        if (? is null or ? = '', f_post.p_desc like '%', f_post.p_desc regexp ?)
      and
        if (? is null or ? = '', f_post.p_body like '%', f_post.p_body regexp ?)
      and
        if(?, not exists (select 1 from ban_post where ban_post.p_id = f_post.p_id), true) # 文章没被封禁
      and
        if(?, not exists (select 1 from ban_user where ban_user.u_id = f_post.u_id), true) # 文章作者没被封禁
      and
        if(?, not exists (select 1 from private_post where private_post.p_id = f_post.p_id), true) # 不是私有文章
    order by ut desc limit ?, ?
) b on a.p_id = b.p_id
order by ut desc;
`
)

func findPost(titleExp string, descExp string, bodyExp string,
	exceptFoldedComment bool, exceptBanComment bool, exceptBanCommentAuthor bool,
	exceptBanPost bool, exceptBanPostAuthor bool, exceptPrivatePost bool,
	sizeOfPage int64, pageIdx int64) (*[]dto.PostWithProp, error) {

	pp := new([]dto.PostWithProp)
	offset := sizeOfPage * pageIdx

	err := dao.QueryN(pp, sql,
		exceptFoldedComment, exceptBanComment, exceptBanCommentAuthor,
		exceptBanComment, exceptBanCommentAuthor, exceptFoldedComment,
		titleExp, titleExp, titleExp,
		descExp, descExp, descExp,
		bodyExp, bodyExp, bodyExp,
		exceptBanPost, exceptBanPostAuthor, exceptPrivatePost, offset, sizeOfPage)
	return pp, err
}

func findPostNoBody(titleExp string, descExp string, bodyExp string,
	exceptFoldedComment bool, exceptBanComment bool, exceptBanCommentAuthor bool,
	exceptBanPost bool, exceptBanPostAuthor bool, exceptPrivatePost bool,
	sizeOfPage int64, pageIdx int64) (*[]dto.PostNoBodyWithProp, error) {

	pp := new([]dto.PostNoBodyWithProp)
	offset := sizeOfPage * pageIdx

	err := dao.QueryN(pp, sql,
		exceptFoldedComment, exceptBanComment, exceptBanCommentAuthor,
		exceptBanComment, exceptBanCommentAuthor, exceptFoldedComment,
		titleExp, titleExp, titleExp,
		descExp, descExp, descExp,
		bodyExp, bodyExp, bodyExp,
		exceptBanPost, exceptBanPostAuthor, exceptPrivatePost, offset, sizeOfPage)
	return pp, err
}

func addTagsAndCategories4Post(tx *sqlx.Tx, postID int64, tagIDs []int64, categoryIDs []int64) error {
	tagSql := `insert into tag_post(p_id, t_id) 
				select ?, ? where not exists (select 1 from tag_post where p_id = ? and t_id = ?)`
	categorySql := `insert into category_post(p_id, cg_id) 
				select ?, ? where not exists (select 1 from category_post where p_id = ? and cg_id = ?)`
	for _, tagID := range tagIDs {
		_, err := tx.Exec(tagSql, postID, tagID, postID, tagID)
		if err != nil {
			return err
		}
	}
	for _, categoryID := range categoryIDs {
		_, err := tx.Exec(categorySql, postID, categoryID, postID, categoryID)
		if err != nil {
			return err
		}
	}
	return nil
}

func add(tx *sqlx.Tx, title, desc, body *string, authorID int64, keywords *string, postType *string, thumbnails *string) (int64, error) {
	sql := `
insert into post(p_title, p_desc, p_body, u_id, p_keywords, p_type, p_thumbnails, p_create_time, p_update_time) 
values(
   ifnull(?, ''), 
   ifnull(?, ''), 
   ifnull(?, ''), 
   ?, 
   ifnull(?, ''), 
   ifnull(?, ''), 
   ifnull(?, ''), now(), now()
)`
	res, err := tx.Exec(sql, title, desc, body, authorID, keywords, postType, thumbnails)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func getByID(p *dto.PostWithProp, id int64) error {
	sql := `
select post.*,
       (select count(like_post.lp_id) from like_post where like_post.p_id = post.p_id and not like_post.lp_neg) as p_like_cnt, 
       (select count(like_post.lp_id) from like_post where like_post.p_id = post.p_id and like_post.lp_neg) as p_unlike_cnt, 
       (select count(comment.c_id) from comment where comment.to_p_id = post.p_id) as p_comment_cnt, 
       (select count(star_post.sp_id) from star_post where star_post.to_p_id = post.p_id) as p_star_cnt
from post where post.p_id = ?
	`
	err := dao.Query1(p, sql, id)
	if err != nil {
		return err
	}

	tp := new([]dto.Tag)
	sql = `select tag.* from tag_post, tag where tag_post.t_id = tag.t_id and p_id = ?`
	err = dao.QueryN(tp, sql, id)
	if err != nil {
		return err
	}

	cp := new([]dto.Category)
	sql = `select category.* from category_post, category where category_post.t_id = category.t_id and p_id = ?`
	err = dao.QueryN(cp, sql, id)
	if err != nil {
		return err
	}

	p.Tags = *tp
	p.Categories = *cp
	return err
}

func getNoBodyByID(p *dto.PostNoBodyWithProp, id int64) error {
	sql := `
select post.*,
       (select count(like_post.lp_id) from like_post where like_post.p_id = post.p_id and not like_post.lp_neg) as p_like_cnt, 
       (select count(like_post.lp_id) from like_post where like_post.p_id = post.p_id and like_post.lp_neg) as p_unlike_cnt, 
       (select count(comment.c_id) from comment where comment.to_p_id = post.p_id) as p_comment_cnt, 
       (select count(star_post.sp_id) from star_post where star_post.to_p_id = post.p_id) as p_star_cnt
from post where post.p_id = ?
	`
	err := dao.Query1(p, sql, id)
	if err != nil {
		return err
	}

	tp := new([]dto.Tag)
	sql = `select tag.* from tag_post, tag where tag_post.t_id = tag.t_id and p_id = ?`
	err = dao.QueryN(tp, sql, id)
	if err != nil {
		return err
	}

	cp := new([]dto.Category)
	sql = `select category.* from category_post, category where category_post.t_id = category.t_id and p_id = ?`
	err = dao.QueryN(cp, sql, id)
	if err != nil {
		return err
	}

	p.Tags = *tp
	p.Categories = *cp
	return err
}

func getAllByUID(uid int64) (*[]dto.PostWithProp, error) {
	sql := `
select post.*,
       (select count(like_post.lp_id) from like_post where like_post.p_id = post.p_id and not like_post.lp_neg) as p_like_cnt,
       (select count(like_post.lp_id) from like_post where like_post.p_id = post.p_id and like_post.lp_neg) as p_unlike_cnt,
       (select count(comment.c_id) from comment where comment.to_p_id = post.p_id) as p_comment_cnt,
       (select count(star_post.sp_id) from star_post where star_post.to_p_id = post.p_id) as p_star_cnt
from post where post.u_id = ?
	`
	pp := new([]dto.PostWithProp)
	err := dao.QueryN(pp, sql, uid)
	return pp, err
}

func getAllNoBodyByUID(uid int64) (*[]dto.PostNoBodyWithProp, error) {
	sql := `
select post.*,
       (select count(like_post.lp_id) from like_post where like_post.p_id = post.p_id and not like_post.lp_neg) as p_like_cnt,
       (select count(like_post.lp_id) from like_post where like_post.p_id = post.p_id and like_post.lp_neg) as p_unlike_cnt,
       (select count(comment.c_id) from comment where comment.to_p_id = post.p_id) as p_comment_cnt,
       (select count(star_post.sp_id) from star_post where star_post.to_p_id = post.p_id) as p_star_cnt
from post where post.u_id = ?
	`
	pp := new([]dto.PostNoBodyWithProp)
	err := dao.QueryN(pp, sql, uid)
	return pp, err
}

func deleteByID(pid int64) (int64, error) {
	sqls := []string{
		`delete from post where p_id = ?`,
		`delete from private_post where p_id = ?`,
		`delete from like_post where p_id = ?`,
		`delete from category_post where p_id = ?`,
		`delete from ban_post where p_id = ?`,
		`delete from star_post where to_p_id = ?`,
		`delete from tag_post where p_id = ?`,
	}
	var lastInsertID int64
	if tx, err := dao.GetTx(); err != nil {
		return 0, err
	} else {
		for _, sql := range sqls {
			res, err := tx.Exec(sql, pid)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
			lastInsertID, _ = res.LastInsertId()
		}
		err := tx.Commit()
		if err != nil {
			return 0, err
		}
	}
	return lastInsertID, nil
}

func updateInfo(p *dto.PostInfoUpdate) (int64, error) {
	sql := `update post set u_email = ?, u_uname = ?, u_pwd_encrypted = sha1(concat(?, 'jun990527')), u_desc = ?, 
           u_thumbnails = ?, u_sex = ?, u_birth = ?, u_tel = ?, u_active_time = now(), u_role = ? where u_id = ?`

	dexss.SimpleText(p.Title)
	dexss.SimpleText(p.Desc)
	dexss.RichText(p.Body) // Body -> RichText
	dexss.SimpleText(p.Keywords)
	dexss.SimpleText(p.Type)
	// 其他字段一般不会被发回，所以无需防御 XSS

	return dao.Update(sql, p.Title, p.Desc, p.AuthorID, p.Keywords, p.Type, p.Thumbnails, p.Body, p.PIDReadOnly)

	// TODO
}
