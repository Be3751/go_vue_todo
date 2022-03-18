# Node.jsの公式イメージ
FROM node:16.14.0-alpine3.15

# Vueプロジェクトとなるディレクトリ
RUN mkdir client

# ワーキングディレクトリの指定
WORKDIR /client

# カレントディレクトリ以下のファイルをコンテナのclientディレクトリにコピー
ADD . /client

CMD ["npm", "run", "serve"]