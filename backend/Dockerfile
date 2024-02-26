FROM golang:1.22-alpine AS builder

WORKDIR /src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /bin .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /bin ./

CMD ["app"]
