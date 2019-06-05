from flask import Flask, jsonify, request
import logging
import MeCab
from gensim.models import KeyedVectors
import pickle
from classifier import tokenize, includeWordList, categorize
from category import sub_categories_ja


logging.basicConfig(format='%(asctime)s : %(levelname)s : %(message)s', level=logging.INFO)

app = Flask(__name__)
app.config['JSON_AS_ASCII'] = False
mecab = MeCab.Tagger("-Ochasen")
model_dir = './model/entity_vector/entity_vector.model.bin'
logging.info("...load word2vec model")
model = KeyedVectors.load_word2vec_format(model_dir, binary=True)

with open('./sub_categories_ja.pickle', mode='rb') as f:
  category_vectors = pickle.load(f)

category_ids = list(category_vectors.keys())
vectors = list(category_vectors.values())


@app.route('/categorize')
def get():
  taglist = request.args.getlist('tag_name')
  tag_category_ids_list = []
  for tag in taglist:
    tag_category_ids = {}
    nornlist = tokenize(mecab, tag)
    logging.info("########## tokenize result #############") 
    logging.info("tokenize result: {}".format(nornlist)) 
    tag_category_ids["tag_name"] = tag
    categoryResults = [categorize(model, category_ids, vectors, norn) for norn in nornlist]
    tag_category_ids["categoryIds"] = categoryResults
    logging.info("########## categorize result #############") 
    logging.info("categorize result: {}".format([sub_categories_ja[i] for i in categoryResults if i != None]))
    tag_category_ids_list.append(tag_category_ids)
   
  return jsonify(tag_category_ids_list)

if __name__ == '__main__':
  app.run(host='0.0.0.0', port=5000)