#! /bin/bash python
from gensim.models import KeyedVectors
import pickle
from category import main_categories

model_dir = './model/entity_vector/entity_vector.model.bin'
model = KeyedVectors.load_word2vec_format(model_dir, binary=True)

vector = {}
for k, v in sub_categories_ja.items():
    vector[k] = model.get_vector(v)

with open('./sub_categories_ja.pickle', mode='wb') as f:
    pickle.dump(vector, f)

