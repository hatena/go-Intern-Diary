package instance

import "github.com/hatena/go-Intern-Diary/model"

func GetInstance() []*model.Category {
	return categoiresInstance
}

var categoiresInstance = []*model.Category{
	&model.Category{1, "スポーツ"},
	&model.Category{2, "フード"},
	&model.Category{3, "技術"},
	&model.Category{4, "アイドル"},

	&model.Category{5, "サッカー"},
	&model.Category{6, "野球"},
	&model.Category{7, "アメフト"},

	&model.Category{8, "お酒"},
	&model.Category{9, "料理"},
	&model.Category{10, "グルメ"},

	&model.Category{11, "Machine Learning"},
	&model.Category{12, "Pyhton"},
	&model.Category{13, "Kubernetes"},

	&model.Category{14, "AKB48"},
	&model.Category{15, "ももクロ"},

	&model.Category{16, "本"},
	&model.Category{17, "日常"},
	&model.Category{18, "子供"},
	&model.Category{19, "アウトドア"},
	&model.Category{20, "写真"},
}
