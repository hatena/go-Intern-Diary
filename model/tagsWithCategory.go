package model

type TagWithCategoriesJson struct {
	TagName     string `json:"tag_name"`
	CategoryIDs []int  `json:"categoryIds`
}

type TagWithCategoryInput struct {
	TagName    string
	CategoryID int32
}

type TagWithCategory struct {
	TagName     string
	CategoryIDs []int
}

func GetTagNamesFromInput(t []*TagWithCategoryInput) []string {
	tag_names := make([]string, len(t))
	for i, tag := range t {
		tag_names[i] = tag.TagName
	}
	return tag_names
}

func ConvertFromInput(t []*TagWithCategoriesJson) []*TagWithCategory {
	newList := make([]*TagWithCategory, 0, len(t))
	for _, tagWithCategory := range t {
		ids := removeNoneCategory(tagWithCategory.CategoryIDs)
		newList = append(newList,
			&TagWithCategory{
				TagName:     tagWithCategory.TagName,
				CategoryIDs: ids,
			},
		)
	}
	return newList
}

func removeNoneCategory(list []int) []int {
	newList := make([]int, 0, len(list))
	for _, x := range list {
		if x > 0 {
			newList = append(newList, x)
		}
	}
	return newList
}
