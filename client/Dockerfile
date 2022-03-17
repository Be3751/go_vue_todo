# Node.jsの公式イメージ
FROM node:16.14.0-alpine3.15

RUN mkdir client

# ワーキングディレクトリの指定
WORKDIR /client

ADD . /client

CMD ["npm", "run", "serve"]