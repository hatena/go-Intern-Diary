package instance

import "github.com/hatena/go-Intern-Diary/model"

func GetInstance() []*model.Category {
	return append(categoiresInstance, subCategoriesInstance...)
}

var categoiresInstance = []*model.Category{
	&model.Category{1, "Entertainment"},
	&model.Category{2, "Showbiz"},
	&model.Category{3, "Sports"},
	&model.Category{4, "Games"},

	&model.Category{5, "Ainmation And Comics"},
	&model.Category{6, "Hobbies"},
	&model.Category{7, "Computers"},

	&model.Category{8, "Gourmet"},
	&model.Category{9, "Style"},
	&model.Category{10, "Life"},

	&model.Category{11, "Leaning And Culture"},
	&model.Category{12, "Regional"},
	&model.Category{13, "Etc"},
}

var subCategoriesInstance = []*model.Category{
	&model.Category{101, "Music"},
	&model.Category{102, "Movies"},
	&model.Category{103, "Japanese Music"},
	&model.Category{104, "Non-Japanese Music"},
	&model.Category{105, "Japanese Movies"},
	&model.Category{106, "Non-Japanese Movies"},

	&model.Category{201, "Idols"},
	&model.Category{202, "Celebrities"},
	&model.Category{203, "Television"},
	&model.Category{204, "Johnnys"},
	&model.Category{205, "AKB48"},
	&model.Category{206, "Comedy"},
	&model.Category{207, "Musicians"},

	&model.Category{301, "Soccer"},
	&model.Category{302, "Basketball"},
	&model.Category{303, "Volleyball"},
	&model.Category{304, "Martial Arts"},
	&model.Category{305, "Golf"},
	&model.Category{306, "Tennis"},
	&model.Category{307, "Other Sports"},

	&model.Category{401, "Game Titles"},
	&model.Category{402, "Game Hardware"},
	&model.Category{403, "Game Companies"},
	&model.Category{404, "Table Games"},
	&model.Category{405, "Flipnote"},
	&model.Category{406, "Animal Crossing"},
	&model.Category{407, "Other Games"},

	&model.Category{501, "Vocaloid"},
	&model.Category{502, "Animation"},
	&model.Category{503, "Comics"},
	&model.Category{504, "Voice Actors"},
	&model.Category{505, "Fanzines"},
	&model.Category{506, "Yaoi"},

	&model.Category{601, "Books"},
	&model.Category{602, "Travel"},
	&model.Category{603, "Photography"},
	&model.Category{604, "Pets"},
	&model.Category{605, "Plants"},
	&model.Category{606, "Automobiles"},
	&model.Category{607, "Bicycles"},
	&model.Category{608, "Motorcycles"},
	&model.Category{609, "Railroads"},
	&model.Category{610, "DIY"},
	&model.Category{611, "Illustration"},
	&model.Category{612, "Other Hobbies"},

	&model.Category{701, "Internet"},
	&model.Category{702, "Tech"},
	&model.Category{703, "Gagdets"},
	&model.Category{704, "Hatena"},
	&model.Category{705, "Programming"},
	&model.Category{706, "Design"},

	&model.Category{801, "Cooking"},
	&model.Category{802, "Restaurants"},
	&model.Category{803, "Alcohol"},
	&model.Category{804, "Food"},
	&model.Category{805, "Drink"},
	&model.Category{806, "Recipes"},
	&model.Category{807, "Sweets"},
	&model.Category{808, "Packed Lunches"},

	&model.Category{901, "Fashion"},
	&model.Category{902, "Beauty"},
	&model.Category{903, "Cosmetics"},
	&model.Category{904, "Mens Fashion"},
	&model.Category{905, "Ladies Fashion"},

	&model.Category{1001, "Childcare"},
	&model.Category{1002, "Families"},
	&model.Category{1003, "Health"},
	&model.Category{1004, "Lifestyles"},
	&model.Category{1005, "Interior Accessories"},
	&model.Category{1006, "Other Life"},

	&model.Category{1101, "Society"},
	&model.Category{1102, "Knowledge"},
	&model.Category{1103, "History"},
	&model.Category{1104, "Other Learning And Culture"},

	&model.Category{1201, "Areas"},
	&model.Category{1202, "Spots"},
	&model.Category{1203, "Shops"},
	&model.Category{1204, "Events"},

	&model.Category{1301, "Chat"},
	&model.Category{1302, "Age Groups"},
	&model.Category{1303, "Personal Attributes"},
	&model.Category{1304, "Fun"},
	&model.Category{1305, "Other Etc"},
}
