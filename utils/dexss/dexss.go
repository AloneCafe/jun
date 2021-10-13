package dexss

import (
	"github.com/microcosm-cc/bluemonday"
)

var (
	ugc    = bluemonday.UGCPolicy()
	strict = bluemonday.StrictPolicy()
)

func SimpleText(ss ...*string) {
	for _, p := range ss {
		if p == nil {
			continue
		}
		*p = strict.Sanitize(*p)
	}
}

func RichText(ss ...*string) {
	for _, p := range ss {
		if p == nil {
			continue
		}
		*p = ugc.Sanitize(*p)
	}
}
