package post

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"jun/dao"
	"jun/dto"
	"jun/model/categroy"
	"jun/model/tag"
)

func FindPost(titleExp string, descExp string, bodyExp string,
	sizeOfPage int64, pageIdx int64) (*[]dto.PostWithProp, error) {

	return findPost(titleExp, descExp, bodyExp,
		true, true, true,
		true, true, true,
		sizeOfPage, pageIdx)
}

func FindPostNoBody(titleExp string, descExp string, bodyExp string,
	sizeOfPage int64, pageIdx int64) (*[]dto.PostNoBodyWithProp, error) {

	return findPostNoBody(titleExp, descExp, bodyExp,
		true, true, true,
		true, true, true,
		sizeOfPage, pageIdx)
}

func Add(title, desc, body *string,
	authorID int64, keywords *string, tagIDs []int64, categoryIDs []int64,
	postType *string, thumbnails *string) (int64, error) {

	if b, err := tag.ExistTagIDs(tagIDs); err != nil {
		// 内部错误
		return 0, err
	} else if !b {
		// 标签不存在
		return 0, errors.New("给文章添加的标签不存在")
	}

	if b, err := category.ExistCategoryIDs(categoryIDs); err != nil {
		// 内部错误
		return 0, err
	} else if !b {
		// 分类不存在
		return 0, errors.New("给文章添加的分类不存在")
	}

	tx, err := dao.GetTx()
	if err != nil {
		return 0, err
	}

	lastInsertID, err := add(tx, title, desc, body, authorID, keywords, postType, thumbnails)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := addTagsAndCategories4Post(tx, lastInsertID, tagIDs, categoryIDs); err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
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
	sql := `insert into post(p_title, p_desc, p_body, u_id, p_keywords, p_type, p_thumbnails) values(?, ?, ?, ?, ?, ?, ?)`
	res, err := tx.Exec(sql, title, desc, body, authorID, keywords, postType, thumbnails)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
