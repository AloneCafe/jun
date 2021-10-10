package post

import (
	"jun/dto"
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

func Add(title string, desc string, body string,
	authorID int64, keywords string, tagIDs []int64, categoryIDs []int64,
	postType string, thumbnails []byte) (int64, error) {

}
