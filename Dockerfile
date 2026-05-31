FROM golang:1.24-alpine

RUN apk add --no-cache git curl

WORKDIR /app

# air (ホットリロード) をインストール
RUN go install github.com/air-verse/air@v1.61.7

EXPOSE 8080

# ソースは docker-compose のボリュームマウントで /app に展開されます
CMD ["air"]
