package post

import (
	"jun/dao"
	"jun/dto"
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

func findPost(titleExp string, descExp string, bodyExp string,
	exceptFoldedComment bool, exceptBanComment bool, exceptBanCommentAuthor bool,
	exceptBanPost bool, exceptBanPostAuthor bool, exceptPrivatePost bool,
	sizeOfPage int64, pageIdx int64) (*[]dto.PostDetail, error) {

	pp := new([]dto.PostDetail)
	offset := sizeOfPage * pageIdx
	sql :=
		`
set @p_title_pattern = trim(?);
set @p_desc_pattern = trim(?);
set @p_body_pattern = trim(?);
set @except_folded_comment = ?;
set @except_ban_comment = ?;
set @except_ban_comment_author= ?;
set @except_ban_post = ?;
set @except_ban_post_author = ?;
set @except_private_post= ?;
set @off = ?;
set @p_size_of_page = ?;
select
    a.*,
    (select count(like_post.lp_id) from like_post where like_post.p_id = a.p_id and not like_post.lp_neg) as p_like_cnt, # 正赞
    (select count(like_post.lp_id) from like_post where like_post.p_id = a.p_id and like_post.lp_neg) as p_unlike_cnt,    # 负赞
    (select count(comment.c_id) from comment where comment.to_p_id = a.p_id
                                               and if(@except_folded_comment, not comment.c_folded, true)
                                               and if(@except_ban_comment, not exists (select 1 from ban_comment where ban_comment.c_id = comment.c_id), true)
                                               and if(@except_ban_comment_author, not exists (select 1 from ban_user, comment where comment.u_id = ban_user.u_id), true)
    ) as p_comment_cnt,    # 评论（看传参是否限制判定）
    (select count(star_post.sp_id) from star_post where star_post.to_p_id = a.p_id) as p_star_cnt # 收藏数
from post a inner join (
    select f_post.p_id,
           if(
                   (exists (select 1 from comment, post p where p.p_id = comment.to_p_id)
                       and if(@except_ban_comment, not exists (select 1 from ban_comment, comment where comment.to_p_id = f_post.p_id and comment.c_id = ban_comment.c_id), true)
                       and if(@except_ban_comment_author, not exists (select 1 from ban_user, comment where comment.to_p_id = f_post.p_id and comment.u_id = ban_user.u_id), true)
                       and if(@except_folded_comment, not (select comment.c_folded from comment where comment.to_p_id = f_post.p_id), true)
                       ), # 考虑评论被封禁、折叠的情况，如果被封禁，就不应该计算其时间
                   (select comment.c_create_time from comment, post p where p.p_id = comment.to_p_id),
                   (select p.p_update_time from post p where p.p_id = f_post.p_id)
               ) as ut from post f_post
    where
        if (@p_title_pattern is null or @p_title_pattern = '', f_post.p_title like '%', f_post.p_title regexp @p_title_pattern)
      and
        if (@p_desc_pattern is null or @p_desc_pattern = '', f_post.p_desc like '%', f_post.p_desc regexp @p_desc_pattern)
      and
        if (@p_body_pattern is null or @p_body_pattern = '', f_post.p_body like '%', f_post.p_body regexp @p_body_pattern)
      and
        if(@except_ban_post, not exists (select 1 from ban_post where ban_post.p_id = f_post.p_id), true) # 文章没被封禁
      and
        if(@except_ban_post_author, not exists (select 1 from ban_user where ban_user.u_id = f_post.u_id), true) # 文章作者没被封禁
      and
        if(@except_private_post, not exists (select 1 from private_post where private_post.p_id = f_post.p_id), true) # 不是私有文章
    order by ut desc limit @off, @p_size_of_page
) b on a.p_id = b.p_id
order by ut desc;
`
	sql =
		`
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
	err := dao.QueryN(pp, sql,
		exceptFoldedComment, exceptBanComment, exceptBanCommentAuthor,
		exceptBanComment, exceptBanCommentAuthor, exceptFoldedComment,
		titleExp, titleExp, titleExp,
		descExp, descExp, descExp,
		bodyExp, bodyExp, bodyExp,
		exceptBanPost, exceptBanPostAuthor, exceptPrivatePost, offset, sizeOfPage)
	return pp, err
}
