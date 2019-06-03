from flask import Flask, jsonify
import logging
import MeCab
from gensim.models import KeyedVectors


logging.basicConfig(format='%(asctime)s : %(levelname)s : %(message)s', level=logging.INFO)

app = Flask(__name__)
app.config['JSON_AS_ASCII'] = False
mecacb = MeCab.Tagger("-Ochasen")
model_dir = './model/entity_vector/entity_vector.model.bin'
model = KeyedVectors.load_word2vec_format(model_dir, binary=True)

with open('./category_vectors', mode='rb') as f:
  category_vactors = pickle.load(f)


@app.route('/')
def index():
  return jsonify({
    "message": "テスト!!"
  })

if __name__ == '__main__':
  app.run(host='0.0.0.0', port=5000)