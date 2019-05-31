package model

type TagWithCategoryInput struct {
	TagName    string
	CategoryID int32
}

type TagWithCategory struct {
	TagName    string
	CategoryID int
}

func ConvertFromInput(t []*TagWithCategoryInput) []*TagWithCategory {
	newList := make([]*TagWithCategory, 0, len(t))
	for _, tagWithCategory := range t {
		newList = append(newList,
			&TagWithCategory{
				TagName:    tagWithCategory.TagName,
				CategoryID: int(tagWithCategory.CategoryID),
			},
		)
	}
	return newList
}
