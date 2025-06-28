# QuillDeck ディレクトリ構成

```
quilldeck/
├── README.md                    # プロジェクト概要・セットアップ手順
├── Makefile                     # ビルド・開発タスク
├── .env.example                 # 環境変数サンプル
├── go.mod                       # Go モジュール定義
├── go.sum                       # Go 依存関係ハッシュ
├── main.go                      # アプリケーションエントリーポイント
│
├── cmd/                         # コマンドライン実行ファイル
│   └── server/                  # サーバー起動コマンド
│       └── main.go
│
├── internal/                    # 内部パッケージ（外部からimport不可）
│   ├── config/                  # 設定管理
│   │   └── config.go
│   │
│   ├── domain/                  # ドメインモデル・ビジネスロジック
│   │   ├── models/              # データモデル定義
│   │   │   ├── user.go
│   │   │   ├── document.go
│   │   │   ├── summary.go
│   │   │   ├── image.go
│   │   │   └── slide.go
│   │   │
│   │   ├── repository/          # データアクセス層インターフェース
│   │   │   ├── user.go
│   │   │   ├── document.go
│   │   │   └── artifact.go
│   │   │
│   │   └── service/             # ビジネスロジック
│   │       ├── auth.go
│   │       ├── document.go
│   │       ├── processing.go
│   │       └── artifact.go
│   │
│   ├── infrastructure/          # 外部システム連携
│   │   ├── database/            # データベース実装
│   │   │   ├── sqlite/
│   │   │   │   ├── connection.go
│   │   │   │   ├── user.go
│   │   │   │   ├── document.go
│   │   │   │   └── artifact.go
│   │   │   └── migrations/      # DBマイグレーション
│   │   │       ├── 001_create_users.sql
│   │   │       ├── 002_create_documents.sql
│   │   │       └── 003_create_artifacts.sql
│   │   │
│   │   ├── ai/                  # AI/MLサービス連携
│   │   │   ├── openai/
│   │   │   │   ├── client.go
│   │   │   │   ├── summarizer.go
│   │   │   │   └── image_generator.go
│   │   │   └── mcp/             # MCP連携
│   │   │       ├── client.go
│   │   │       └── protocol.go
│   │   │
│   │   ├── storage/             # ファイルストレージ
│   │   │   ├── local.go
│   │   │   └── uploader.go
│   │   │
│   │   └── export/              # エクスポート機能
│   │       ├── pdf.go
│   │       └── pptx.go
│   │
│   ├── interfaces/              # 外部インターフェース層
│   │   ├── http/                # HTTP API
│   │   │   ├── server.go
│   │   │   ├── middleware/
│   │   │   │   ├── auth.go
│   │   │   │   ├── cors.go
│   │   │   │   └── logging.go
│   │   │   │
│   │   │   ├── handlers/        # HTTPハンドラー
│   │   │   │   ├── auth.go
│   │   │   │   ├── document.go
│   │   │   │   ├── processing.go
│   │   │   │   └── artifact.go
│   │   │   │
│   │   │   └── dto/             # データ転送オブジェクト
│   │   │       ├── auth.go
│   │   │       ├── document.go
│   │   │       └── processing.go
│   │   │
│   │   └── cli/                 # CLI インターフェース（将来拡張用）
│   │       └── commands.go
│   │
│   └── pkg/                     # 内部共通ユーティリティ
│       ├── logger/              # ログ機能
│       │   └── logger.go
│       ├── validator/           # バリデーション
│       │   └── validator.go
│       ├── crypto/              # 暗号化ユーティリティ
│       │   └── hash.go
│       └── errors/              # エラーハンドリング
│           └── errors.go
│
├── web/                         # フロントエンド
│   ├── package.json
│   ├── tsconfig.json
│   ├── next.config.js
│   │
│   ├── src/
│   │   ├── components/          # Reactコンポーネント
│   │   │   ├── common/          # 共通コンポーネント
│   │   │   │   ├── Layout.tsx
│   │   │   │   ├── Header.tsx
│   │   │   │   └── Loading.tsx
│   │   │   │
│   │   │   ├── auth/            # 認証関連
│   │   │   │   ├── LoginForm.tsx
│   │   │   │   └── RegisterForm.tsx
│   │   │   │
│   │   │   ├── document/        # ドキュメント関連
│   │   │   │   ├── FileUpload.tsx
│   │   │   │   ├── DocumentList.tsx
│   │   │   │   └── DocumentViewer.tsx
│   │   │   │
│   │   │   ├── processing/      # 処理関連
│   │   │   │   ├── SummaryGenerator.tsx
│   │   │   │   ├── ImageGenerator.tsx
│   │   │   │   └── SlideGenerator.tsx
│   │   │   │
│   │   │   └── artifact/        # 成果物関連
│   │   │       ├── ArtifactList.tsx
│   │   │       ├── ArtifactPreview.tsx
│   │   │       └── ArtifactDownload.tsx
│   │   │
│   │   ├── pages/               # Next.jsページ
│   │   │   ├── _app.tsx
│   │   │   ├── _document.tsx
│   │   │   ├── index.tsx
│   │   │   ├── login.tsx
│   │   │   ├── dashboard.tsx
│   │   │   └── artifacts/
│   │   │       └── [id].tsx
│   │   │
│   │   ├── hooks/               # カスタムフック
│   │   │   ├── useAuth.ts
│   │   │   ├── useDocument.ts
│   │   │   └── useArtifact.ts
│   │   │
│   │   ├── services/            # API通信
│   │   │   ├── api.ts
│   │   │   ├── auth.ts
│   │   │   ├── document.ts
│   │   │   └── processing.ts
│   │   │
│   │   ├── types/               # TypeScript型定義
│   │   │   ├── auth.ts
│   │   │   ├── document.ts
│   │   │   └── artifact.ts
│   │   │
│   │   └── utils/               # ユーティリティ
│   │       ├── constants.ts
│   │       ├── helpers.ts
│   │       └── validation.ts
│   │
│   └── public/                  # 静的ファイル
│       ├── favicon.ico
│       └── images/
│
├── storage/                     # ローカルストレージ
│   ├── uploads/                 # アップロードファイル
│   ├── generated/               # 生成ファイル
│   │   ├── summaries/
│   │   ├── images/
│   │   └── slides/
│   └── exports/                 # エクスポートファイル
│
├── configs/                     # 設定ファイル
│   ├── database.yaml
│   ├── ai_services.yaml
│   └── server.yaml
│
├── scripts/                     # 開発・運用スクリプト
│   ├── setup.sh
│   ├── migrate.sh
│   └── build.sh
│
├── docs/                        # ドキュメント
│   ├── api/                     # API仕様書
│   ├── architecture/            # アーキテクチャ設計書
│   └── deployment/              # デプロイメント手順
│
└── tests/                       # テストファイル
    ├── unit/                    # 単体テスト
    ├── integration/             # 結合テスト
    └── e2e/                     # E2Eテスト
```

## 設計思想

### 1. Clean Architecture採用
- `domain/`: ビジネスロジックを中心に配置
- `infrastructure/`: 外部システム依存を分離
- `interfaces/`: 外部からのアクセス点を管理

### 2. 機能別モジュール分割
- 認証、ドキュメント、処理、成果物管理を独立したモジュールとして構成
- 各機能の責務を明確に分離

### 3. 拡張性を考慮
- MCP連携、AI/MLサービス追加に対応
- データベース切り替え可能な設計
- フロントエンド・バックエンド独立開発

### 4. 開発効率重視
- 型安全性（TypeScript/Go）
- テスト容易性（依存注入、インターフェース分離）
- 設定外部化（YAML/環境変数）