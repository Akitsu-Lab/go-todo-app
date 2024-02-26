FROM golang:1.22-alpine AS builder

WORKDIR /src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# バイナリを /bin/app にビルドする
RUN go build -v -o /app .

FROM alpine:latest

# バイナリをコピーして実行可能にする
COPY --from=builder /app /usr/local/bin/app

CMD ["app"]
