from gensim.models import KeyedVectors
import gensim
import pickle

NORN_WORD_LIST = [
    '名詞-一般',
    '名詞-固有'
]

# return norn list
def tokenize(mecab, input):
    lineList = mecab.parse(input).split('\n')
    nornList = [line.split('\t')[0] for line in lineList if '\t' in line and includeWordList(line.split('\t')[3], NORN_WORD_LIST)]
    return nornList


def includeWordList(word, wordList):
    isInclude = [ w in word for w in wordList ]
    return sum(isInclude) > 0
    
def categorize(model, category_vectors, word):
    vector = model.get_vector(word)
    ids, vectors = category_vectors.items()
    maxIndex = cosine_similarities(vector, vectors).argmax()
    return ids[maxIndex]

