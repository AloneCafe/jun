package option

import (
	"strconv"
)

func GetWebsiteName() (string, error) {
	wn, err := getOption("OPTION_WEBSITE")
	if err != nil {
		return "", err
	} else {
		return *wn, err
	}
}

func GetPostCountPerPage() (int64, error) {
	pcpp, err := getOption("OPTION_POST_COUNT_PER_PAGE")
	if err != nil {
		return 0, err
	}
	c, err := strconv.ParseInt(*pcpp, 10, 64)
	if err != nil {
		return 0, err
	}
	return c, nil
}
