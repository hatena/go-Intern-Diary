FROM python:3.6

ARG project_dir=/word2vec/
WORKDIR $project_dir

RUN apt-get update \
  && apt-get install -y mecab \
  && apt-get install -y mecab-ipadic \
  && apt-get install -y libmecab-dev \
  && apt-get install -y mecab-ipadic-utf8 \
  && apt-get install -y swig

COPY . $project_dir

RUN pip install -r requirements.txt