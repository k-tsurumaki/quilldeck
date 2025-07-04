# Changelog

## [v0.2.0] - 2025-06-28

### Added
- **フロントエンド実装**
  - React + TypeScript + Vite環境構築
  - 認証フォーム（登録・ログイン切り替え）
  - ファイルアップロード機能
  - 要約生成インターフェース
  - レスポンシブデザイン対応

- **CORS対応**
  - 全APIエンドポイントにCORSヘッダー追加
  - フロントエンド・バックエンド間通信対応

- **UI/UX改善**
  - エラーハンドリング強化
  - ローディング状態表示
  - ユーザーフレンドリーなメッセージ

### Fixed
- Backend connection failedエラーの修正
- ファイルアップロード失敗エラーの修正
- APIクライアントの絶対URL対応

### Changed
- Docker Compose設定の最適化
- README.mdの大幅更新

## [v0.1.0] - 2025-06-28

### Added
- **バックエンド基盤**
  - Go + fuselageフレームワーク
  - SQLiteデータベース連携
  - ユーザー認証API（登録・ログイン）
  - ファイルアップロードAPI（TXT/MD対応）
  - 要約生成API（サンプル実装）

- **インフラ構築**
  - Docker環境構築
  - マルチステージビルド
  - データベースマイグレーション
  - ヘルスチェックエンドポイント

- **ドメイン設計**
  - クリーンアーキテクチャ採用
  - リポジトリパターン実装
  - エラーハンドリング統一

### Technical Details
- **技術スタック**: Go 1.23, SQLite, Docker
- **フレームワーク**: fuselage (HTTP), database/sql
- **アーキテクチャ**: Clean Architecture
- **テスト**: 基本的なAPI動作確認済み