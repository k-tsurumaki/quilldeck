.PHONY: test test-unit test-coverage clean deps

# テスト実行
test:
	go test -v ./...

# 単体テストのみ実行
test-unit:
	go test -v ./tests/unit/...

# カバレッジ付きテスト実行
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# 依存関係のダウンロード
deps:
	go mod download
	go mod tidy

# クリーンアップ
clean:
	go clean
	rm -f coverage.out coverage.html

# ビルド
build:
	go build -o bin/quilldeck ./cmd/server

# 開発用サーバー起動
dev:
	go run ./cmd/server