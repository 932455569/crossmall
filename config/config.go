package config

import "net/http"

// Config CROSS MALL 基础配置
type Config struct {
	// Account アカウント (必须)
	// 对应 CROSSMALL管理画面＞システム設定＞基本情報＞会社コード
	Account string

	// AuthKey 認証鍵 (必须)
	// 对应 API設定画面で設定した認証鍵
	AuthKey string

	// BaseURL ドメイン名 (可选)
	// 默认 "https://crossmall.jp"
	BaseURL string

	// HTTPClient 自定义 HTTP 客户端 (可选)
	// 默认使用 http.DefaultClient
	HTTPClient *http.Client
}