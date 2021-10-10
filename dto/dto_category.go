package dto

type Category struct {
	ID   int64   `json:"t_id" db:"t_id"`
	Name *string `json:"t_name" db:"t_name"`
}

func DetachCategoriesIDs(ts []Category) []int64 {
	var r []int64
	for _, t := range ts {
		r = append(r, t.ID)
	}
	return r
}

func DetachCategoriesNames(ts []Category) []string {
	var r []string
	for _, t := range ts {
		r = append(r, *t.Name)
	}
	return r
}
