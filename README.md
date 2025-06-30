<div align="center">

# 🪶 QuillDeck

**AI-Powered Document Summarization Platform**

[![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![React](https://img.shields.io/badge/React-18+-61DAFB?style=for-the-badge&logo=react&logoColor=black)](https://reactjs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5+-3178C6?style=for-the-badge&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)

[🇺🇸 English](#-english) • [🇯🇵 日本語](#-japanese)

</div>

---

## 🇺🇸 English

### ✨ Overview

QuillDeck transforms your documents into actionable insights using cutting-edge AI technology. Upload, analyze, and extract key information in seconds.

### 🚀 Features

<table>
<tr>
<td align="center">📄</td>
<td><strong>Smart Upload</strong><br/>Support for TXT and MD files with drag & drop</td>
</tr>
<tr>
<td align="center">🤖</td>
<td><strong>AI Summarization</strong><br/>Powered by advanced LLM APIs for accurate summaries</td>
</tr>
<tr>
<td align="center">🔐</td>
<td><strong>Secure Auth</strong><br/>User management with encrypted authentication</td>
</tr>
<tr>
<td align="center">🎨</td>
<td><strong>Modern UI</strong><br/>Responsive design with multi-language support</td>
</tr>
</table>

### ⚡ Quick Start

```bash
# 📥 Clone & Setup
git clone <repository-url>
cd quilldeck
cp .env.example .env

# 🔧 Configure your LLM API keys in .env

# 🚀 Launch
docker compose up -d
```

<div align="center">

**🌐 Access Points**

[Frontend →](http://localhost:3000) `localhost:3000` | [API →](http://localhost:8080) `localhost:8080`

</div>

### 🏗️ Tech Stack

| Layer | Technology | Purpose |
|-------|------------|----------|
| **Frontend** | React + TypeScript + Tailwind | Modern, responsive UI |
| **Backend** | Go + Fuselage | High-performance API |
| **Database** | SQLite / PostgreSQL | Reliable data storage |
| **AI** | OpenRouter API | Advanced summarization |

### 📡 API Reference

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/auth/register` | `POST` | 👤 User registration |
| `/api/auth/login` | `POST` | 🔑 User authentication |
| `/api/documents/upload` | `POST` | 📤 Document upload |
| `/api/documents/summary` | `POST` | 🤖 Generate summary |

### 🛠️ Development

```bash
# 🧪 Run tests
go test ./...

# 🔨 Build backend
go build ./cmd/server

# 💻 Frontend dev
cd web && npm install && npm run dev
```

---

## 🇯🇵 Japanese

### ✨ 概要

QuillDeckは最先端のAI技術を使用して、ドキュメントを実用的な洞察に変換します。アップロード、分析、重要な情報の抽出を数秒で実現。

### 🚀 機能

<table>
<tr>
<td align="center">📄</td>
<td><strong>スマートアップロード</strong><br/>ドラッグ&ドロップ対応のTXT・MDファイルサポート</td>
</tr>
<tr>
<td align="center">🤖</td>
<td><strong>AI要約</strong><br/>高度なLLM APIによる正確な要約生成</td>
</tr>
<tr>
<td align="center">🔐</td>
<td><strong>セキュア認証</strong><br/>暗号化認証によるユーザー管理</td>
</tr>
<tr>
<td align="center">🎨</td>
<td><strong>モダンUI</strong><br/>多言語対応のレスポンシブデザイン</td>
</tr>
</table>

### ⚡ クイックスタート

```bash
# 📥 クローン & セットアップ
git clone <repository-url>
cd quilldeck
cp .env.example .env

# 🔧 .envファイルでLLM APIキーを設定

# 🚀 起動
docker compose up -d
```

<div align="center">

**🌐 アクセスポイント**

[フロントエンド →](http://localhost:3000) `localhost:3000` | [API →](http://localhost:8080) `localhost:8080`

</div>

### 🏗️ 技術スタック

| レイヤー | 技術 | 用途 |
|----------|------|------|
| **フロントエンド** | React + TypeScript + Tailwind | モダンでレスポンシブなUI |
| **バックエンド** | Go + Fuselage | 高性能API |
| **データベース** | SQLite / PostgreSQL | 信頼性の高いデータストレージ |
| **AI** | OpenRouter API | 高度な要約機能 |

### 📡 API リファレンス

| エンドポイント | メソッド | 説明 |
|----------------|----------|------|
| `/api/auth/register` | `POST` | 👤 ユーザー登録 |
| `/api/auth/login` | `POST` | 🔑 ユーザー認証 |
| `/api/documents/upload` | `POST` | 📤 ドキュメントアップロード |
| `/api/documents/summary` | `POST` | 🤖 要約生成 |

### 🛠️ 開発

```bash
# 🧪 テスト実行
go test ./...

# 🔨 バックエンドビルド
go build ./cmd/server

# 💻 フロントエンド開発
cd web && npm install && npm run dev
```

---

<div align="center">

**Made with ❤️ using Go & React**

*Transform your documents into insights*

</div>