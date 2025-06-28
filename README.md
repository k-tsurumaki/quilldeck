# QuillDeck

QuillDeckは、ドキュメントのアップロードと要約生成を行うWebアプリケーションです。

## 🏗️ アーキテクチャ

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │    Backend      │    │   Database      │
│   (React)       │◄──►│   (Go/fuselage) │◄──►│   (SQLite)      │
│   Port: 3000    │    │   Port: 8080    │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### 技術スタック
- **Backend**: Go + fuselage フレームワーク
- **Frontend**: React + TypeScript + Vite
- **Database**: SQLite (開発環境)
- **Infrastructure**: Docker + Docker Compose

## 🚀 クイックスタート

### 前提条件
- Docker
- Docker Compose

### 起動方法

1. **全サービス起動**
```bash
docker compose up -d
```

2. **アクセス**
- フロントエンド: http://localhost:3000
- バックエンドAPI: http://localhost:8080

3. **使用方法**
- ブラウザで http://localhost:3000 にアクセス
- ユーザー登録またはログイン
- TXTまたはMDファイルをアップロード
- 要約を生成

4. **APIテスト（オプション）**
```bash
# ヘルスチェック
curl http://localhost:8080/health

# ユーザー登録
curl -X POST -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123","name":"Test User"}' \
  http://localhost:8080/api/auth/register

# ファイルアップロード
curl -X POST -F "file=@your_file.txt" \
  http://localhost:8080/api/documents/upload
```

## 📁 プロジェクト構造

```
quilldeck/
├── cmd/server/           # アプリケーションエントリーポイント
├── internal/
│   ├── config/          # 設定管理
│   ├── domain/          # ドメインロジック
│   │   ├── models/      # エンティティ
│   │   ├── repository/  # リポジトリインターフェース
│   │   └── service/     # ビジネスロジック
│   ├── infrastructure/  # インフラ層
│   │   └── database/    # データベース実装
│   └── interfaces/      # プレゼンテーション層
│       └── http/        # HTTPハンドラー
├── web/                 # フロントエンド
├── data/                # SQLiteデータベースファイル
├── docker-compose.yml   # Docker構成
└── Dockerfile          # Dockerイメージ定義
```

## 🔧 開発環境

### ローカル開発（Dockerなし）

1. **依存関係インストール**
```bash
go mod tidy
```

2. **データディレクトリ作成**
```bash
mkdir -p data
```

3. **アプリケーション起動**
```bash
go run cmd/server/main.go
```

### フロントエンド開発

```bash
cd web
npm install
npm run dev
```

## 📋 API仕様

### 認証
- `POST /api/auth/register` - ユーザー登録
- `POST /api/auth/login` - ログイン

### ドキュメント
- `POST /api/documents/upload` - ファイルアップロード
- `POST /api/documents/summary` - 要約生成

### システム
- `GET /health` - ヘルスチェック

## ✅ 実装済み機能

### バックエンド
- [x] fuselageフレームワークによるHTTPサーバー
- [x] SQLiteデータベース連携
- [x] ユーザー認証（登録・ログイン）
- [x] ファイルアップロード（TXT/MD対応）
- [x] 要約生成（サンプル実装：一文目抽出）
- [x] CORS対応
- [x] Docker環境構築
- [x] データベースマイグレーション

### フロントエンド
- [x] React + TypeScript + Vite環境
- [x] 認証フォーム（登録・ログイン切り替え）
- [x] ファイルアップロード機能
- [x] 要約生成インターフェース
- [x] レスポンシブデザイン
- [x] エラーハンドリング
- [x] ローディング状態管理

### インフラ
- [x] Docker Compose設定
- [x] マルチステージビルド
- [x] SQLite永続化
- [x] フロントエンド・バックエンド連携

## 📝 TODO（今後の実装予定）

### 高優先度
- [ ] **AI連携による要約生成**
  - OpenAI API統合
  - MCP (Model Context Protocol) 対応
  - 要約品質の向上

- [ ] **認証・認可の強化**
  - JWT トークン実装
  - セッション管理
  - パスワードハッシュ化

- [ ] **UI/UXの改善**
  - ダークモード対応
  - ファイルドラッグ&ドロップ
  - プログレスバー表示

### 中優先度
- [ ] **ファイル処理の拡張**
  - PDF対応
  - Word文書対応
  - ファイルサイズ制限設定

- [ ] **データ管理機能**
  - ドキュメント一覧表示
  - 要約履歴管理
  - ファイル削除機能
  - 検索機能

- [ ] **データベース拡張**
  - PostgreSQL対応
  - データベース接続プール
  - マイグレーション管理

### 低優先度
- [ ] **運用・監視**
  - ログ管理
  - メトリクス収集
  - ヘルスチェック拡張

- [ ] **セキュリティ**
  - CORS設定
  - レート制限
  - 入力値検証強化

- [ ] **パフォーマンス**
  - キャッシュ機能
  - 非同期処理
  - ファイル圧縮

## 🐛 トラブルシューティング

### よくある問題

**フロントエンドが表示されない**
```bash
# サービス状態確認
docker compose ps
# フロントエンドログ確認
docker compose logs frontend
```

**Backend connection failed エラー**
```bash
# バックエンド状態確認
curl http://localhost:8080/health
# サービス再起動
docker compose restart backend
```

**ポート8080が使用中**
```bash
# 使用中のプロセスを確認
lsof -i :8080
# プロセスを停止
kill <PID>
```

**Docker ビルドエラー**
```bash
# キャッシュクリア
docker system prune -a
# 再ビルド
docker compose build --no-cache
```

## 📄 ライセンス

MIT License

## 🤝 コントリビューション

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request