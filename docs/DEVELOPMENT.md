# 開発者ガイド

## Docker環境の詳細

### サービス構成

```yaml
services:
  backend:    # Go アプリケーション
  frontend:   # React アプリケーション  
  database:   # PostgreSQL (将来用)
```

### ネットワーク構成

```
quilldeck-network (bridge)
├── backend:8080
├── frontend:3000  
└── database:5432
```

### データ永続化

- `./data:/app/data` - SQLiteデータベース
- `postgres_data` - PostgreSQLボリューム

## 開発ワークフロー

### 1. バックエンド開発

```bash
# 依存関係更新
go mod tidy

# テスト実行
go test ./...

# ローカル起動
go run cmd/server/main.go

# Docker再ビルド
docker compose build backend
```

### 2. フロントエンド開発

```bash
cd web
npm install
npm run dev    # 開発サーバー
npm run build  # プロダクションビルド
```

### 3. データベース操作

```bash
# SQLiteファイル確認
sqlite3 data/quilldeck.db ".tables"

# データ確認
sqlite3 data/quilldeck.db "SELECT * FROM users;"
```

## API テスト例

### 完全なワークフロー

```bash
# 1. ユーザー登録
USER_RESPONSE=$(curl -s -X POST -H "Content-Type: application/json" \
  -d '{"email":"dev@example.com","password":"dev123","name":"Developer"}' \
  http://localhost:8080/api/auth/register)

echo $USER_RESPONSE

# 2. ファイルアップロード
DOC_RESPONSE=$(curl -s -X POST -F "file=@test_document.txt" \
  http://localhost:8080/api/documents/upload)

echo $DOC_RESPONSE

# 3. ドキュメントIDを抽出して要約生成
DOCUMENT_ID=$(echo $DOC_RESPONSE | jq -r '.document_id')

curl -s -X POST -H "Content-Type: application/json" \
  -d "{\"document_id\":\"$DOCUMENT_ID\",\"length\":\"short\"}" \
  http://localhost:8080/api/documents/summary
```

## デバッグ

### ログ確認

```bash
# アプリケーションログ
tail -f server.log

# Dockerログ
docker compose logs -f backend
```

### データベースデバッグ

```bash
# テーブル構造確認
sqlite3 data/quilldeck.db ".schema users"

# データ投入テスト
sqlite3 data/quilldeck.db "INSERT INTO users (id, email, name, password) VALUES ('test-id', 'test@example.com', 'Test', 'password');"
```