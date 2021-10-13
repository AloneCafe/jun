package post

import (
	"errors"

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

func GetNoBodyByID(id int64) (*dto.PostNoBodyWithProp, error) {
	p := new(dto.PostNoBodyWithProp)
	err := getNoBodyByID(p, id)
	return p, err
}

func GetByID(id int64) (*dto.PostWithProp, error) {
	p := new(dto.PostWithProp)
	err := getByID(p, id)
	return p, err
}

func GetAllByUID(uid int64) (*[]dto.PostWithProp, error) {
	return getAllByUID(uid)
}

func GetAllNoBodyByUID(uid int64) (*[]dto.PostNoBodyWithProp, error) {
	return getAllNoBodyByUID(uid)
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

func DeleteByID(pid int64) (int64, error) {
	return deleteByID(pid)
}

func UpdateInfo(p *dto.PostInfoUpdate) (int64, error) {
	return updateInfo(p)
}
