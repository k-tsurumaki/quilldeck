FROM golang:1.23.6-alpine AS builder

# Cコンパイラを追加
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

# 依存関係をコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# CGO有効でビルド（SQLite用）
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o quilldeck ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates sqlite
WORKDIR /root/

# バイナリをコピー
COPY --from=builder /app/quilldeck .

# データディレクトリ作成
RUN mkdir -p /root/data

EXPOSE 8080

CMD ["./quilldeck"]