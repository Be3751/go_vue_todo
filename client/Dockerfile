# Node.jsの公式イメージ
FROM node:16.14.0-alpine3.15

# ワーキングディレクトリの指定
WORKDIR /src

# Vueのインストール
RUN apk update && \
    npm install && \
    npm install -g npm && \
    npm install -g vue-cli