# Goの公式イメージ
FROM golang:1.17.8-alpine3.15

# docker run時に作業ディレクトリの作成とgitのインストールを実行
RUN mkdir /go/src/server
RUN apk update && apk add git

# ワーキングディレクトリの指定
WORKDIR /go/src/server

# ホストのファイル・ディレクトリをコンテナイメージにコピー
ADD . /go/src/server 