package model

type Category struct {
	ID           uint64 `db:"id"`
	CategoryName string `db:"category_name"`
}

type CategoryNode struct {
	ID           int
	CategoryName string
	ParentNode   *CategoryNode
}

var CategoryRoot = &CategoryNode{0, "root", nil}

var Sports = &CategoryNode{1, "スポーツ", CategoryRoot}
var Food = &CategoryNode{2, "フード", CategoryRoot}
var Technics = &CategoryNode{3, "技術", CategoryRoot}
var Idle = &CategoryNode{4, "アイドル", CategoryRoot}

var Category5 = &CategoryNode{5, "サッカー", Sports}
var Category6 = &CategoryNode{6, "野球", Sports}
var Category7 = &CategoryNode{7, "アメフト", Sports}

var Category8 = &CategoryNode{8, "お酒", Food}
var Category9 = &CategoryNode{9, "料理", Food}
var Category10 = &CategoryNode{10, "グルメ", Food}

var Category11 = &CategoryNode{11, "Machine Learning", Technics}
var Category12 = &CategoryNode{12, "Pyhton", Technics}
var Category13 = &CategoryNode{13, "Kubernetes", Technics}

var Category14 = &CategoryNode{14, "AKB48", Idle}
var Category15 = &CategoryNode{15, "ももクロ", Idle}
