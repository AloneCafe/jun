package post

import (
	"jun/dto"
)

func FindPost(titleExp string, descExp string, bodyExp string,
	sizeOfPage int64, pageIdx int64) (*[]dto.PostDetail, error) {

	return findPost(titleExp, descExp, bodyExp,
		true, true, true,
		true, true, true,
		sizeOfPage, pageIdx)
}
