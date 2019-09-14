# CASystem
CASystemの個人開発用リポジトリ

## システム構成
参考文献(https://qiita.com/tozastation/items/a69a102fdc3f62d566b4)

```
├── pkg:
│   ├── application: ビジネスロジック
│   ├── di: 依存性の注入
│   ├── infra: 外部との通信
│   └── server
│       ├── handler: エンドポイント
│       ├── middleware: 認証
│       ├── response
│       └── server.go: ルーティング
└── main.go: 
```

## 技術選定

- 技術選定
    - アプリケーション: Go
        - アーキテクチャ: Layered Architecture
        - ログ: zap
        - 認証・認可: Firebase
        - ルーティング: net
        - マイグレーション: goose
        - データベースアクセス: gorm
        - テスト: Go test
        - モック・スタブ: Go mock
    - データベース: MySQL
