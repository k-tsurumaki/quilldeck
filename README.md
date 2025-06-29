# QuillDeck

[English](#english) | [日本語](#日本語)

---

## English

### Overview
QuillDeck is an AI-powered document summarization platform that helps users quickly extract key insights from their documents.

### Features
- **Document Upload**: Support for TXT and MD files
- **AI Summarization**: Automatic content summarization using LLM APIs
- **User Management**: Secure authentication and user accounts
- **Modern UI**: Clean, responsive web interface

### Quick Start

#### Prerequisites
- Docker and Docker Compose
- Git

#### Installation
```bash
# Clone the repository
git clone <repository-url>
cd quilldeck

# Copy environment configuration
cp .env.example .env

# Edit .env file with your settings
# Set your LLM API keys and other configurations

# Start the services
docker compose up -d
```

#### Access
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

### Architecture
- **Frontend**: React with TypeScript and Tailwind CSS
- **Backend**: Go with Echo framework
- **Database**: SQLite (development) / PostgreSQL (production)
- **AI Integration**: OpenRouter API support

### API Endpoints
- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User login
- `POST /api/documents/upload` - Upload document
- `POST /api/documents/summary` - Generate summary

### Development
```bash
# Run tests
go test ./...

# Build backend
go build ./cmd/server

# Frontend development
cd web
npm install
npm run dev
```

### Environment Variables
See `.env.example` for all available configuration options.

---

## 日本語

### 概要
QuillDeckは、ユーザーがドキュメントから重要な洞察を素早く抽出できるAI搭載の文書要約プラットフォームです。

### 機能
- **ドキュメントアップロード**: TXTおよびMDファイルのサポート
- **AI要約**: LLM APIを使用した自動コンテンツ要約
- **ユーザー管理**: 安全な認証とユーザーアカウント
- **モダンUI**: クリーンでレスポンシブなWebインターフェース

### クイックスタート

#### 前提条件
- Docker と Docker Compose
- Git

#### インストール
```bash
# リポジトリをクローン
git clone <repository-url>
cd quilldeck

# 環境設定をコピー
cp .env.example .env

# .envファイルを編集して設定を行う
# LLM APIキーやその他の設定を行ってください

# サービスを開始
docker compose up -d
```

#### アクセス
- フロントエンド: http://localhost:3000
- バックエンドAPI: http://localhost:8080

### アーキテクチャ
- **フロントエンド**: React with TypeScript and Tailwind CSS
- **バックエンド**: Go with Echo framework
- **データベース**: SQLite (開発環境) / PostgreSQL (本番環境)
- **AI統合**: OpenRouter API サポート

### APIエンドポイント
- `POST /api/auth/register` - ユーザー登録
- `POST /api/auth/login` - ユーザーログイン
- `POST /api/documents/upload` - ドキュメントアップロード
- `POST /api/documents/summary` - 要約生成

### 開発
```bash
# テスト実行
go test ./...

# バックエンドビルド
go build ./cmd/server

# フロントエンド開発
cd web
npm install
npm run dev
```

### 環境変数
利用可能な設定オプションについては `.env.example` を参照してください。