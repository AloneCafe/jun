package tag

func ExistTagIDs(ids []int64) (bool, error) {
	return existTagIDs(ids)
}

func ExistTagNames(names []string) (bool, error) {
	return existTagNames(names)
}
