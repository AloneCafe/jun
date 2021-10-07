package post

import (
	"jun/dto"
)

func FindPost(titleExp string, descExp string, bodyExp string,
	sizeOfPage int, pageIdx int) (*[]dto.PostDetail, error) {

	return findPost(titleExp, descExp, bodyExp,
		true, true, true,
		true, true, true,
		sizeOfPage, pageIdx)
}
