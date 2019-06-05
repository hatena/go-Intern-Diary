from gensim.models import KeyedVectors
import gensim
import pickle
from app import logging

NORN_WORD_LIST = [
    '名詞-一般',
    '名詞-固有'
]

# return nornlist :[str]
def tokenize(mecab, input):
    lineList = mecab.parse(input).split('\n')
    nornList = [line.split('\t')[0] for line in lineList if '\t' in line and includeWordList(line.split('\t')[3], NORN_WORD_LIST)]
    return nornList


# return :bool
def includeWordList(word, wordList):
    isInclude = [ w in word for w in wordList ]
    return sum(isInclude) > 0
    

# return category_id: int
def categorize(model, category_ids, vectors, word):
    try:    
        vector = model.get_vector(word)
        maxIndex = model.cosine_similarities(vector, vectors).argmax()
        return category_ids[maxIndex]
    except KeyError:
        logging.warning("word {} is not in vocabulary")

