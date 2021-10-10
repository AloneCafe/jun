package category

func ExistCategoryIDs(ids []int64) (bool, error) {
	return existCategoryIDs(ids)
}

func ExistCategoryNames(names []string) (bool, error) {
	return existCategoryNames(names)
}
