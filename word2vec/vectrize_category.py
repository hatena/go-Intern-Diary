#! /bin/bash python
from gensim.models import KeyedVectors
import pickle

model_dir = './model/entity_vector/entity_vector.model.bin'
model = KeyedVectors.load_word2vec_format(model_dir, binary=True)
sub_categories = {

    101: "Music",
	102: "Movies",
	103: "Japanese Music",
	104: "Non-Japanese Music",
	105: "Japanese Movies",
	106: "Non-Japanese Movies",

	201: "Idols",
	202: "Celebrities",
	203: "Television",
	204: "Johnnys",
	205: "AKB48",
	206: "Comedy",
	207: "Musicians",

	301: "Soccer",
	302: "Basketball",
	303: "Volleyball",
	304: "Martial Arts",
	305: "Golf",
	306: "Tennis",
	307: "Other Sports",

	401: "Game Titles",
	402: "Game Hardware",
	403: "Game Companies",
	404: "Table Games",
	405: "Flipnote",
	406: "Animal Crossing",
	407: "Other Games",

	501: "Vocaloid",
	502: "Animation",
	503: "Comics",
	504: "Voice Actors",
	505: "Fanzines",
	506: "Yaoi",

	601: "Books",
	602: "Travel",
	603: "Photography",
	604: "Pets",
	605: "Plants",
	606: "Automobiles",
	607: "Bicycles",
	608: "Motorcycles",
	609: "Railroads",
	610: "DIY",
	611: "Illustration",
	612: "Other Hobbies",

	701: "Internet",
	702: "Tech",
	703: "Gagdets",
	704: "Hatena",
	705: "Programming",
	706: "Design",

	801: "Cooking",
	802: "Restaurants",
	803: "Alcohol",
	804: "Food",
	805: "Drink",
	806: "Recipes",
	807: "Sweets",
	808: "Packed Lunches",

	901: "Fashion",
	902: "Beauty",
	903: "Cosmetics",
	904: "Mens Fashion",
	905: "Ladies Fashion",

	1001: "Childcare",
	1002: "Families",
	1003: "Health",
	1004: "Lifestyles",
	1005: "Interior Accessories",
	1006: "Other Life",

	1101: "Society",
	1102: "Knowledge",
	1103: "History",
	1104: "Other Learning And Culture",

	1201: "Areas",
	1202: "Spots",
	1203: "Shops",
	1204: "Events",

	1301: "Chat",
	1302: "Age Groups",
	1303: "Personal Attributes",
	1304: "Fun",
	1305: "Other Etc"
}

main_categories = {
    1: "Entertainment",
	2: "ショービズ",
	3: "Sports",
	4: "Games",

	# 5: "Animation And Comics",
    5: "Animation",
    14: "Comics",
	6: "Hobbies",
	7: "Computer",

	8: "Gourmet",
	9: "Style",
	10: "Life",

	# 11: "Leaning And Culture",
    11: "Learning",
    15: "Culture",
	12: "Regional",
	13: "Etc"
}

vector = {}
for k, v in main_categories.items():
    vector[k] = model.get_vector(v)

with open('./category_vectors', mode='wb') as f:
    pickle.dump(vector, f)

